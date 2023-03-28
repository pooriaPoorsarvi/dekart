package apiBqjob

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	"dekart/src/proto"
	"dekart/src/server/job"
	"dekart/src/server/storage"

	bigquery "google.golang.org/api/bigquery/v2"

	"google.golang.org/api/googleapi"
)

// Job implements the dekart.Job interface for BigQuery; concurrency safe.
type Job struct {
	job.BasicJob
	bigqueryJob         *bigquery.Job
	storageObject       storage.StorageObject
	maxReadStreamsCount int32
	maxBytesBilled      int64
	client 				*bigquery.Service
	token 				*oauth2.Token
}

var contextCancelledRe = regexp.MustCompile(`context canceled`)
var orderByRe = regexp.MustCompile(`(?ims)order[\s]+by`)

func (job *Job) close(storageWriter io.WriteCloser, csvWriter *csv.Writer) {
	csvWriter.Flush()
	err := storageWriter.Close()
	if err != nil {
		// maybe we should not close when context is canceled in job.write()
		if err == context.Canceled {
			return
		}
		if contextCancelledRe.MatchString(err.Error()) {
			return
		}
		job.Logger.Err(err).Send()
		job.CancelWithError(err)
		return
	}
	resultSize, err := job.storageObject.GetSize(job.GetCtx())
	if err != nil {
		job.Logger.Err(err).Send()
		job.CancelWithError(err)
		return
	}

	job.Logger.Debug().Msg("Writing Done")
	job.Lock()
	job.ResultSize = *resultSize
	jobID := job.GetID()
	job.ResultID = &jobID
	job.Unlock()
	job.Status() <- int32(proto.Query_JOB_STATUS_DONE)
	job.Cancel()
}

func (job *Job) setJobStats(queryStatus *bigquery.Job, table *bigquery.Table) error {
	if table == nil {
		return fmt.Errorf("table is nil")
	}


	job.Lock()
	defer job.Unlock()
	if queryStatus.Statistics != nil {
		job.ProcessedBytes = queryStatus.Statistics.TotalBytesProcessed
	}
	job.TotalRows = int64(table.NumRows)
	return nil
}

// write csv rows to storage
func (job *Job) write(csvRows chan []string) {
	storageWriter := job.storageObject.GetWriter(job.GetCtx())
	csvWriter := csv.NewWriter(storageWriter)
	for {
		csvRow, more := <-csvRows
		if !more {
			break
		}
		err := csvWriter.Write(csvRow)
		if err == context.Canceled {
			break
		}
		if err != nil {
			job.Logger.Err(err).Send()
			job.CancelWithError(err)
			break
		}
	}
	job.close(storageWriter, csvWriter)
}

type AvroSchema struct {
	Fields []struct {
		Name string `json:"name"`
	} `json:"fields"`
}

func (job *Job) processApiErrors(err error) {
	if apiError, ok := err.(*googleapi.Error); ok {
		for _, e := range apiError.Errors {
			if e.Reason == "bytesBilledLimitExceeded" {
				job.Logger.Warn().Str(
					"DEKART_BIGQUERY_MAX_BYTES_BILLED", os.Getenv("DEKART_BIGQUERY_MAX_BYTES_BILLED"),
				).Msg(e.Message)
			}
		}
	}
}

func (job *Job) getResultTable() (*bigquery.Table, error) {
	srcTbl := job.bigqueryJob.Configuration.Query.DestinationTable
	table, err := job.client.Tables.Get(srcTbl.ProjectId, srcTbl.DatasetId, srcTbl.TableId).Do()
	if err != nil {
		job.Logger.Err(err)
		return nil, err
	}
	return table, nil
}

//func (job *Job) GetResultTableForScript() (*bigquery.Table, error){
//
//
//
//	jobFromJobId, err := job.client.JobFromID(job.GetCtx(), job.bigqueryJob.ID())
//	if err != nil{
//		return nil, err
//	}
//
//	cfg, err := jobFromJobId.Config()
//
//	if err != nil{
//		return nil, err
//	}
//
//	queryConfig, ok := cfg.(*bigquery.QueryConfig)
//	if !ok{
//		err := fmt.Errorf("was expecting QueryConfig type for configuration")
//		job.Logger.Error().Err(err).Str("jobConfig", fmt.Sprintf("%v+", cfg)).Send()
//		return nil, err
//	}
//
//	table := queryConfig.Dst
//
//	if table == nil {
//		err := fmt.Errorf("destination table is nil even when gathered from JobId")
//		job.Logger.Error().Err(err).Str("jobConfig", fmt.Sprintf("%v+", cfg)).Send()
//		return nil, err
//	}
//
//	return table, nil
//
//}


func (job *Job) wait() {

	newJobInfo, err := job.client.Jobs.Get(job.bigqueryJob.JobReference.ProjectId, job.bigqueryJob.JobReference.JobId).Do()
	for {
		job.bigqueryJob, err = job.client.Jobs.Get(job.bigqueryJob.JobReference.ProjectId, job.bigqueryJob.JobReference.JobId).Do()
		if err != nil {
			break
		}
		if newJobInfo.Status.State == "DONE" {
			if newJobInfo.Status.ErrorResult != nil {
				err = errors.New(fmt.Sprint("job failed with error: %v", newJobInfo.Status.ErrorResult.Message))
				break
			}
			break
		}
		time.Sleep(5 * time.Second)
	}
	job.bigqueryJob = newJobInfo
	queryStatus := newJobInfo

	if err == context.Canceled {
		return
	}
	if err != nil {
		job.processApiErrors(err)
		job.CancelWithError(err)
		return
	}
	if queryStatus.Status == nil {
		job.Logger.Fatal().Msgf("queryStatus == nil")
	}
	if queryStatus.Status != nil && queryStatus.Status.Errors != nil && len(queryStatus.Status.Errors) > 0 {
		errorsString := []string{}
		for i := 0; i < len(queryStatus.Status.Errors); i++ {
			errorsString = append(errorsString, queryStatus.Status.Errors[i].Reason + " : " + queryStatus.Status.Errors[i].Message)
		}
		job.CancelWithError(errors.New(strings.Join(errorsString, ",")))
		return
	}

	table, err := job.getResultTable()
	if err != nil {
		job.CancelWithError(err)
		return
	}

	err = job.setJobStats(queryStatus, table)
	if err != nil {
		job.CancelWithError(err)
		return
	}

	job.Status() <- int32(proto.Query_JOB_STATUS_READING_RESULTS)

	csvRows := make(chan []string, job.TotalRows)
	errors := make(chan error)

	job.Logger.Debug().Msg("Reading")
	// read table rows into csvRows
	go Read(
		//job.GetCtx(),
		context.Background(),
		job.token,
		errors,
		csvRows,
		table,
		job.Logger,
		job.maxReadStreamsCount,
	)
	job.Logger.Debug().Msg("finished Reading")

	// write csvRows to storage
	go job.write(csvRows)

	// wait for errors
	err = <-errors
	if err != nil {
		job.CancelWithError(err)
		return
	}
	job.Logger.Debug().Msg("Job Wait Done")
}

func (job *Job) setMaxReadStreamsCount(queryText string) {
	job.Lock()
	defer job.Unlock()
	if orderByRe.MatchString(queryText) {
		job.maxReadStreamsCount = 1 // keep order of items
	} else {
		job.maxReadStreamsCount = 10
	}
}

// Run implementation
func (job *Job) Run(storageObject storage.StorageObject) error {
	token := job.GetToken()
	job.token = token
	oauthClient := oauth2.NewClient(job.GetCtx(), oauth2.StaticTokenSource(job.token))
	client, err := bigquery.NewService(job.GetCtx(), option.WithHTTPClient(oauthClient))
	if err != nil {
		job.Cancel()
		return err
	}
	job.client = client

	UseLegacySql := false
	queryRequest := &bigquery.QueryRequest{
		Query: job.QueryText,
		// TODO : Remove maxbytes billed for now, check later
		// MaximumBytesBilled: job.maxBytesBilled,
		UseLegacySql: &UseLegacySql,
	}
	query, err := job.client.Jobs.Query(os.Getenv("DEKART_BIGQUERY_PROJECT_ID"), queryRequest).Do()
	if err != nil{
		job.Logger.Err(err)
		return err
	}

	job.setMaxReadStreamsCount(job.QueryText)

	bigqueryJob, err := job.client.Jobs.Get(query.JobReference.ProjectId, query.JobReference.JobId).Do()
	if err != nil {
		job.Logger.Err(err)
		job.Cancel()
		return err
	}
	job.Lock()
	job.bigqueryJob = bigqueryJob
	job.storageObject = storageObject
	job.Unlock()
	job.Status() <- int32(proto.Query_JOB_STATUS_RUNNING)
	job.Logger.Debug().Msg("Waiting for results")
	go job.wait()
	return nil
}

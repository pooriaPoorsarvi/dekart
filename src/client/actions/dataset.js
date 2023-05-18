import { CreateDatasetRequest, RemoveDatasetRequest } from '../../proto/dekart_pb'
import { Dekart } from '../../proto/dekart_pb_service'
import { unary } from '../lib/grpc'
import { downloading, error, finishDownloading, success, slowDownLimit, alertLimit } from './message'
import {
  addDataToMap,
  toggleSidePanel,
  removeDataset as removeDatasetFromKepler,
  receiveMapConfig,
  loadFiles
} from 'kepler.gl/dist/actions'
import { processCsvData, processGeojson } from 'kepler.gl/dist/processors'
import { get } from '../lib/api'
import { KeplerGlSchema } from 'kepler.gl/dist/schemas'

export function createDataset (reportId) {
  return (dispatch) => {
    dispatch({ type: createDataset.name })
    const request = new CreateDatasetRequest()
    request.setReportId(reportId)
    unary(Dekart.CreateDataset, request).catch(err => dispatch(error(err)))
  }
}

export function setActiveDataset (datasetId) {
  return (dispatch, getState) => {
    const { datasets } = getState()
    const dataset = datasets.find(d => d.id === datasetId) || datasets[0]
    if (dataset) {
      dispatch({ type: setActiveDataset.name, dataset })
    }
  }
}

export function removeDataset (datasetId) {
  return async (dispatch, getState) => {
    const { datasets, activeDataset } = getState()
    if (activeDataset.id === datasetId) {
      // removed active query
      const datasetsLeft = datasets.filter(q => q.id !== datasetId && q.id !== datasetId.split(".")[0])
      if (datasetsLeft.length === 0) {
        dispatch(error(new Error('Cannot remove last dataset')))
        return
      }
      dispatch(setActiveDataset(datasetsLeft.id))
    }
    dispatch({ type: removeDataset.name, datasetId })
    let csvId = datasetId + ".csv";
    dispatch(removeDatasetFromKepler(csvId))
    dispatch(removeDatasetFromKepler(datasetId))


    const request = new RemoveDatasetRequest()
    request.setDatasetId(datasetId)
    try {
      await unary(Dekart.RemoveDataset, request)
      dispatch(success('Dataset removed'))
    } catch (err) {
      dispatch(error(err))
    }
  }
}

export function loadFile (dataset, sourceId, extension, label) {
  return { 
    type: loadFile.name, 
    downloadDatasetValue:{
      "dataset":dataset,
      "sourceId": sourceId, 
      "extension": extension, 
      "label": label
    }
  }
}


export function loadFileStarted (dataset, sourceId, extension, label) {
  return { 
    type: loadFileStarted.name, 
    downloadDatasetValue:{
      "dataset":dataset,
      "sourceId": sourceId, 
      "extension": extension, 
      "label": label
    }
  }
}

const addedDatasets = []

export function downloadDataset (dataset, sourceId, extension, label) {
  return async (dispatch, getState) => {
    dispatch({ type: downloadDataset.name, dataset })
    dispatch(downloading(dataset))
    let data
    try {
      const res = await get(`/dataset-source/${sourceId}.${extension}`)
      if (extension === 'csv') {
        
        let res_clone = res.clone();
        let blob = await res_clone.blob();
        const {fileLoadNeeded} = getState();
        let csv;
        if(fileLoadNeeded.filter(d => d.dataset.id === dataset.id).length === 0){
          csv = await res.text();
          // console.log("csv is");
          // console.log(csv);
          if(blob.size > 0 && csv.length === 0){
            dispatch(finishDownloading(dataset))
            dispatch(loadFile(dataset, sourceId, extension, label))
            return;
          }
          data = processCsvData(csv)
        }else{
          const files = new File([blob], `${dataset.id}.${extension}`)
          dispatch(loadFiles([files]))
          dispatch(loadFileStarted(dataset, sourceId, extension, label))
          dispatch(toggleSidePanel('layer'))
          return;
        }
        
        // if (res.siz)
      } else {
        const json = await res.json()
        data = processGeojson(json)
      }
    } catch (err) {
      dispatch(error(err))
      return
    }
    const { datasets, keplerGl } = getState()
    const i = datasets.findIndex(d => d.id === dataset.id)
    if (i < 0) {
      return
    }
    try {
      if (addedDatasets.includes(dataset.id)) {
        // kepler does not update datasets correctly
        // so we have to remove and add again

        // receive config
        const config = KeplerGlSchema.getConfigToSave(keplerGl.kepler)

        // remove dataset
        dispatch(removeDatasetFromKepler(dataset.id))

        // add dataset with previous config
        dispatch(addDataToMap({
          datasets: {
            info: {
              label,
              id: dataset.id
            },
            data
          },
          config
        }))
      } else {
        dispatch(addDataToMap({
          datasets: {
            info: {
              label,
              id: dataset.id
            },
            data
          }
        }))
        addedDatasets.push(dataset.id)
      }
      
      if (data.rows.length > slowDownLimit){
        dispatch(alertLimit(data.rows.length, dataset.id));
      }
    } catch (err) {
      console.log(err)
      dispatch(error(
        new Error(`Failed to add data to map: ${err.message}`),
        false
      ))
      return
    }
    dispatch(finishDownloading(dataset))
    const { reportStatus } = getState()
    if (reportStatus.edit) {
      dispatch(toggleSidePanel('layer'))
    }
  }
}

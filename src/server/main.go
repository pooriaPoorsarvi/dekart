package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"dekart/src/proto"
	"dekart/src/server/dekart"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func configureLogger() {
	rand.Seed(time.Now().UnixNano())
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	pretty := os.Getenv("DEKART_LOG_PRETTY")
	if pretty != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	debug := os.Getenv("DEKART_LOG_DEBUG")
	if debug != "" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Info().Msgf("Log level: %s", zerolog.GlobalLevel().String())

}

func configureDb() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DEKART_POSTGRES_USER"),
		os.Getenv("DEKART_POSTGRES_PASSWORD"),
		os.Getenv("DEKART_POSTGRES_HOST"),
		os.Getenv("DEKART_POSTGRES_PORT"),
		os.Getenv("DEKART_POSTGRES_DB"),
	))
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)
	return db
}

func applyMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("WithInstance")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal().Err(err).Msg("NewWithDatabaseInstance")
	}
	m.Up()
	if err != nil {
		log.Fatal().Err(err).Msg("Migrations Up")
	}
}

func main() {
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	configureLogger()

	db := configureDb()
	defer db.Close()

	applyMigrations(db)

	dekartServer := dekart.Server{
		Db: db,
	}

	grpcServer := grpc.NewServer()
	proto.RegisterDekartServer(grpcServer, dekartServer)
	grpcwebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			// log.Debug().Str("origin", origin).Send()
			//TODO:
			return true
		}),
	)

	port := os.Getenv("DEKART_PORT")
	log.Info().Msgf("Starting dekart at :%s", port)
	httpServer := &http.Server{
		Handler: http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
			grpcwebServer.ServeHTTP(resp, req)
		}),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal().Err(httpServer.ListenAndServe()).Send()

	// port := os.Getenv("DEKART_PORT")
	// log.Info().Msgf("Starting dekart at :%s", port)
	// listener, err := net.Listen("tcp", "localhost:"+port)
	// if err != nil {
	// 	log.Fatal().Err(err).Send()
	// }

	// server.Serve(listener)
	// reflection.Register(s)
	// r := mux.NewRouter()
	// api := r.PathPrefix("/api/v1").Subrouter()
	// api.Use(mux.CORSMethodMiddleware(r))

	// // POST /v1/api/report
	// api.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 	if r.Method == http.MethodOptions {
	// 		return
	// 	}
	// 	reportsManager.CreateReportHandler(ctx, w, r)
	// }).Methods("POST", "OPTIONS")

	// // POST /v1/api/report/$id/query
	// api.HandleFunc("/report/{reportId}/query", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 	if r.Method == http.MethodOptions {
	// 		return
	// 	}
	// 	vars := mux.Vars(r)
	// 	reportsManager.CreateQueryHandler(ctx, vars["reportId"], w, r)
	// }).Methods("POST", "OPTIONS")

	// // PATCH /v1/api/query/$id
	// api.HandleFunc("/query/{queryId}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 	w.Header().Set("Access-Control-Allow-Methods", "PATCH")
	// 	if r.Method == http.MethodOptions {
	// 		return
	// 	}
	// 	vars := mux.Vars(r)
	// 	reportsManager.UpdateQueryHandler(ctx, vars["queryId"], w, r)
	// }).Methods("PATCH", "OPTIONS")

	// // GET /v1/api/report/$id
	// api.HandleFunc("/report/{reportId}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 	if r.Method == http.MethodOptions {
	// 		return
	// 	}
	// 	vars := mux.Vars(r)
	// 	reportsManager.GetReportHandler(ctx, vars["reportId"], w, r)
	// }).Methods("GET", "OPTIONS")

	// port := os.Getenv("DEKART_PORT")
	// log.Info().Msgf("Starting dekart at :%s", port)
	// srv := &http.Server{
	// 	Handler:      r,
	// 	Addr:         ":" + port,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	// log.Fatal().Err(srv.ListenAndServe()).Send()

}

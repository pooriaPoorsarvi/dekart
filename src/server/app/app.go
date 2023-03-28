package app

import (
	"dekart/src/proto"
	"dekart/src/server/dekart"
	"dekart/src/server/user"
	"encoding/json"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

// ResponseWriter implementation which allows to oweride status code
type ResponseWriter struct {
	w          http.ResponseWriter
	statusCode int
}

// Header implementation
func (m ResponseWriter) Header() http.Header {
	return m.w.Header()
}

// Write implementation
func (m ResponseWriter) Write(b []byte) (int, error) {
	return m.w.Write(b)
}

// WriteHeader overrides statusOk with configured header
func (m ResponseWriter) WriteHeader(statusCode int) {
	if statusCode != http.StatusOK {
		log.Warn().Int("statusCode", statusCode).Send()
		m.w.WriteHeader(statusCode)
	} else {
		m.w.WriteHeader(m.statusCode)
	}
}

var allowedOrigin string = os.Getenv("DEKART_CORS_ORIGIN")

func configureGRPC(dekartServer *dekart.Server) *grpcweb.WrappedGrpcServer {
	server := grpc.NewServer()
	proto.RegisterDekartServer(server, dekartServer)
	return grpcweb.WrapServer(
		server,
		grpcweb.WithOriginFunc(func(origin string) bool {
			if allowedOrigin == "" || allowedOrigin == "*" {
				//log.Warn().Msg("DEKART_CORS_ORIGIN is empty or *")
				return true
			}
			result := origin == allowedOrigin
			if !result {
				//log.Warn().Str("origin", origin).Str("allowed origin", allowedOrigin).Msg("Origin is not allowed")
			}
			return result
		}),
	)
}

func setOriginHeader(w http.ResponseWriter, r *http.Request) {
	if allowedOrigin == "" || allowedOrigin == "*" {
		//log.Warn().Msg("DEKART_CORS_ORIGIN is empty or *")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	}
}

func configureHTTP(dekartServer *dekart.Server) *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1/").Subrouter()
	api.Use(mux.CORSMethodMiddleware(router))

	conf := dekart.OauthConf


	api.HandleFunc("/init-authenticate-oauth2", func(w http.ResponseWriter, r *http.Request) {

		setOriginHeader(w, r)
		if r.Method == http.MethodOptions {
			return
		}

		if ! dekart.ShouldUseTokens(){
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&map[string]interface{}{"redirectUrl": "", "authorizationNeeded": false})
			return
		}

		ctx := r.Context()
		claims := user.GetClaims(ctx)
		_, err := dekartServer.RetrieveToken(claims.Email)

		if err != nil  {
			if err == dekart.RetrievalNotSupported{
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&map[string]interface{}{"redirectUrl": url, "authorizationNeeded": true})
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&map[string]interface{}{"redirectUrl": "", "authorizationNeeded": false})
		}





	}).Methods("POST", "GET", "OPTIONS")


	api.HandleFunc("/callback-authenticate-oauth2", func(w http.ResponseWriter, r *http.Request) {
		setOriginHeader(w, r)

		if r.Method == http.MethodOptions {
			return
		}

		dekartServer.SaveToken(r.URL.Query().Get("code"), r.URL.Query().Get("state"))
		http.Redirect(w, r, "/", http.StatusSeeOther)


	}).Methods("POST", "GET", "OPTIONS")




	api.HandleFunc("/dataset-source/{id}.{extension:csv|geojson}", func(w http.ResponseWriter, r *http.Request) {
		setOriginHeader(w, r)
		if r.Method == http.MethodOptions {
			return
		}
		dekartServer.ServeDatasetSource(w, r)
	}).Methods("GET", "OPTIONS")

	api.HandleFunc("/query-source/{id}.sql", func(w http.ResponseWriter, r *http.Request) {
		setOriginHeader(w, r)
		if r.Method == http.MethodOptions {
			return
		}
		dekartServer.ServeQuerySource(w, r)
	}).Methods("GET", "OPTIONS")

	api.HandleFunc("/file/{id}.csv", func(w http.ResponseWriter, r *http.Request) {
		setOriginHeader(w, r)
		if r.Method == http.MethodOptions {
			return
		}
		dekartServer.UploadFile(w, r)
	}).Methods("POST", "OPTIONS")

	staticPath := os.Getenv("DEKART_STATIC_FILES")

	if staticPath != "" {
		staticFilesHandler := NewStaticFilesHandler(staticPath)

		router.HandleFunc("/", staticFilesHandler.ServeIndex)
		router.HandleFunc("/reports/{id}", staticFilesHandler.ServeIndex)
		router.HandleFunc("/reports/{id}/edit", staticFilesHandler.ServeIndex) // deprecated
		router.HandleFunc("/reports/{id}/source", staticFilesHandler.ServeIndex)
		router.HandleFunc("/400", func(w http.ResponseWriter, r *http.Request) {
			staticFilesHandler.ServeIndex(ResponseWriter{w: w, statusCode: http.StatusBadRequest}, r)
		})
		router.PathPrefix("/").Handler(staticFilesHandler)
	} else {
		log.Warn().Msg("DEKART_STATIC_FILES is empty; UI will not be served")
	}

	return router
}

// Configure HTTP server with http and grpc
func Configure(dekartServer *dekart.Server) *http.Server {
	grpcServer := configureGRPC(dekartServer)
	httpServer := configureHTTP(dekartServer)
	claimsCheck := user.NewClaimsCheck(
		os.Getenv("DEKART_IAP_JWT_AUD"),
		os.Getenv("DEKART_REQUIRE_IAP") == "1",
		os.Getenv("DEKART_REQUIRE_AMAZON_OIDC") == "1",
		os.Getenv("AWS_REGION"),
		os.Getenv("DEKART_DEV_CLAIMS_EMAIL"),
	)

	port := os.Getenv("DEKART_PORT")
	log.Info().Msgf("Starting dekart at :%s", port)
	return &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqWithClaims := r.WithContext(claimsCheck.GetContext(r))
			if grpcServer.IsAcceptableGrpcCorsRequest(r) || grpcServer.IsGrpcWebRequest(r) {
				grpcServer.ServeHTTP(w, reqWithClaims)
			} else {
				httpServer.ServeHTTP(w, reqWithClaims)
			}
		}),
		Addr:         ":" + port,
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
}

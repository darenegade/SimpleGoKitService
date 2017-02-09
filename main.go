package main

import (
	"log"
	"net/http"
	"os"

	stdjwt "github.com/dgrijalva/jwt-go"
	httptransport "github.com/go-kit/kit/transport/http"
	LOG "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/auth/jwt"
	"golang.org/x/net/context"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"

	"github.com/darenegade/SimpleGoKitService/database"
	"github.com/darenegade/SimpleGoKitService/hello"
	"github.com/darenegade/SimpleGoKitService/middleware"
	"github.com/darenegade/SimpleGoKitService/employee"
)

var (
	logger = LOG.NewLogfmtLogger(os.Stdout)
	ctx = context.Background()
	kf = func(token *stdjwt.Token) (interface{}, error) { return []byte("TEST"), nil }
)

func main() {

	database.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT not set")
	}

	r := mux.NewRouter()

	handleHelloWorld(r)
	handleEmployees(r)

	http.Handle("/", accessControl(r))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHelloWorld(r *mux.Router) {

	endpointPath, createdEndpoint := hello.MakeHelloWorldEndpoint()
	createdEndpoint = configEndpoint(createdEndpoint, endpointPath)

	helloWorldHandler := httptransport.NewServer(
		ctx,
		createdEndpoint,
		hello.MakeHelloWorldDecoder(),
		httptransport.EncodeJSONResponse,
		httptransport.ServerBefore(jwt.ToHTTPContext()),
		httptransport.ServerErrorLogger(logger),
	)

	r.Handle(endpointPath, helloWorldHandler).Methods(http.MethodPost)

}

func handleEmployees(r *mux.Router){

	endpointPath, createdEndpoint := employee.MakeEmployeesEndpoint()
	createdEndpoint = configEndpoint(createdEndpoint, endpointPath)

	handler := httptransport.NewServer(
		ctx,
		createdEndpoint,
		employee.MakeEmployeesDecoder(),
		httptransport.EncodeJSONResponse,
		httptransport.ServerBefore(jwt.ToHTTPContext()),
		httptransport.ServerErrorLogger(logger),
	)


	r.Handle(endpointPath, handler).Methods(http.MethodGet, http.MethodPost)

	endpointPath, createdEndpoint = employee.MakeEmployeeEndpoint()
	createdEndpoint = configEndpoint(createdEndpoint, endpointPath)

	handler = httptransport.NewServer(
		ctx,
		createdEndpoint,
		employee.MakeEmployeeDecoder(),
		httptransport.EncodeJSONResponse,
		httptransport.ServerBefore(jwt.ToHTTPContext()),
		httptransport.ServerErrorLogger(logger),
	)

	r.Handle(endpointPath, handler).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
}

func configEndpoint (endpoint endpoint.Endpoint, endpointName string) endpoint.Endpoint {
	endpoint = jwt.NewParser(kf, stdjwt.SigningMethodHS256)(endpoint)
	endpoint = middleware.Logging(LOG.NewContext(logger).With("method", endpointName))(endpoint)
	return endpoint
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
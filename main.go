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

	"github.com/go-kit/kit/endpoint"

	"github.com/darenegade/SimpleGoKitService/database"
	"github.com/darenegade/SimpleGoKitService/hello"
	"github.com/darenegade/SimpleGoKitService/util"
	"github.com/darenegade/SimpleGoKitService/middleware"
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

	handleHelloWorld()


	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHelloWorld() {

	svc := hello.HelloWorldService{}
	endpointPath, helloWorldEndpoint := hello.MakeHelloWorldEndpoint(svc)
	helloWorldEndpoint = configEndpoint(helloWorldEndpoint, endpointPath)

	helloWorldHandler := httptransport.NewServer(
		ctx,
		helloWorldEndpoint,
		hello.MakeHelloWorldDecoder(),
		util.EncodeResponse,
		httptransport.ServerBefore(jwt.ToHTTPContext()),
		httptransport.ServerErrorLogger(logger),
	)
	http.Handle(endpointPath, helloWorldHandler)
}

func configEndpoint (endpoint endpoint.Endpoint, endpointName string) endpoint.Endpoint {
	endpoint = jwt.NewParser(kf, stdjwt.SigningMethodHS256)(endpoint)
	endpoint = middleware.Logging(LOG.NewContext(logger).With("method", endpointName))(endpoint)
	return endpoint
}
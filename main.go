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
)

func main() {

	database.Initialize()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT not set")
	}

	logger := LOG.NewLogfmtLogger(os.Stdout)
	ctx := context.Background()
	kf := func(token *stdjwt.Token) (interface{}, error) { return []byte("TEST"), nil }

	handleHelloWorld(logger, kf, ctx)


	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHelloWorld(logger LOG.Logger, kf func(token *stdjwt.Token) (interface{}, error), ctx context.Context) {
	var svc HelloWorldService
	var helloWorldEndpoint endpoint.Endpoint

	svc = helloWorldService{}
	helloWorldEndpoint = makeHelloWorldEndpoint(svc)
	helloWorldEndpoint = jwt.NewParser(kf, stdjwt.SigningMethodHS256)(helloWorldEndpoint)
	helloWorldEndpoint = Logging(LOG.NewContext(logger).With("method", "hello_service"))(helloWorldEndpoint)
	helloWorldHandler := httptransport.NewServer(
		ctx,
		helloWorldEndpoint,
		decodeHelloWorldRequest,
		encodeResponse,
		httptransport.ServerBefore(jwt.ToHTTPContext()),
		httptransport.ServerErrorLogger(logger),
	)
	http.Handle("/hello_service", helloWorldHandler)
}
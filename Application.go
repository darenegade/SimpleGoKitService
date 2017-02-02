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
)

func main() {
	logger := LOG.NewLogfmtLogger(os.Stdout)
	ctx := context.Background()

	var svc HelloWorldService
	svc = helloWorldService{}
	svc = loggingMiddleware{logger, svc}

	var helloWorldEndpoint endpoint.Endpoint
	helloWorldEndpoint = makeHelloWorldEndpoint(svc)
	kf := func(token *stdjwt.Token) (interface{}, error) { return []byte("TEST"), nil }
	helloWorldEndpoint = jwt.NewParser(kf, stdjwt.SigningMethodHS256)(helloWorldEndpoint)

	helloWorldHandler := httptransport.NewServer(
		ctx,
		helloWorldEndpoint,
		decodeHelloWorldRequest,
		encodeResponse,
		httptransport.ServerBefore(jwt.ToHTTPContext()),
		httptransport.ServerErrorLogger(logger),
	)

	http.Handle("/hello_service", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
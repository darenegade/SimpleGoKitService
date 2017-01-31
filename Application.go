package main

import (
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	LOG "github.com/go-kit/kit/log"
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
	//kf := func(token *stdjwt.Token) (interface{}, error) { return []byte("TEST"), nil }
	//helloWorldEndpoint = jwt.NewParser(kf, stdjwt.SigningMethodHS256)(helloWorldEndpoint)

	helloWorldHandler := httptransport.NewServer(
		ctx,
		helloWorldEndpoint,
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/hello_service", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
package hello

import (
	"github.com/go-kit/kit/endpoint"
	http2 "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"net/http"

	"github.com/darenegade/SimpleGoKitService/util"
)

func MakeHelloWorldEndpoint(svc HelloWorldService) (string, endpoint.Endpoint) {
	return "/hello_service", func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(util.Request)

		if req.Method == http.MethodPost {
			return svc.HelloService(*(req.Data.(*HelloWorld)))
		} else {
			return nil, util.ErrWrongMethod
		}
	}
}

func MakeHelloWorldDecoder() http2.DecodeRequestFunc {
	return util.MakeDecoder(func() interface{} {
		return new(HelloWorld)
	})



}

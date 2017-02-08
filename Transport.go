package main

import (
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"

	"golang.org/x/net/context"
	"errors"
	http2 "github.com/go-kit/kit/transport/http"
)

type ErrorWithStatus struct {
	error
	code int
}

func (error ErrorWithStatus) StatusCode() int { return error.code}

var ErrWrongMethod = ErrorWithStatus{ errors.New("Request has wrong method") , 405}

func makeHelloWorldEndpoint(svc HelloWorldService) (string, endpoint.Endpoint) {
	return "/hello_service", func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)

		if req.Method == http.MethodPost {
			return svc.helloService(*(req.data.(*helloWorld)))
		} else {
			return nil, ErrWrongMethod
		}
	}
}

func makeHelloWorldDecoder() http2.DecodeRequestFunc {
	return makeDecoder(func() interface{} {
		return new(helloWorld)
	})



}

func makeDecoder(data createNewDatatype) http2.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {

		body := data()

		if err := json.NewDecoder(r.Body).Decode(body); err != nil {
			return nil, err
		}

		request := Request{r,body}
		return request, nil
	}
}

type createNewDatatype func() interface{}

type Request struct {
	*http.Request
	data interface{}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}





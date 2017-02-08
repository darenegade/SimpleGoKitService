package main

import (
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"

	"golang.org/x/net/context"
	"errors"
)

type ErrorWithStatus struct {
	error
	code int
}

func (error ErrorWithStatus) StatusCode() int { return error.code}

var ErrWrongMethod = ErrorWithStatus{ errors.New("Request has wrong method") , 405}

func makeHelloWorldEndpoint(svc HelloWorldService) (string, endpoint.Endpoint) {
	return "/hello_service", func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(helloWorldRequest)
		v, err := svc.helloService(req.Name)
		if err != nil {
			return helloWorldResponse{v, err.Error()}, nil
		}
		return helloWorldResponse{v,""}, nil
	}
}

func decodeHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request helloWorldRequest

	if r.Method != http.MethodPost {
		return nil, ErrWrongMethod
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	ResponseText   string `json:"response"`
	Err string `json:"err,omitempty"`
}


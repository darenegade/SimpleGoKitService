package main

import (
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"

	"golang.org/x/net/context"
)

func makeHelloWorldEndpoint(svc HelloWorldService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
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


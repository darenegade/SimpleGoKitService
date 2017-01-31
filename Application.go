package main

import (
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"github.com/go-kit/kit/endpoint"
	"errors"
)

func main() {
	ctx := context.Background()
	svc := helloWorldService{}

	helloWorldHandler := httptransport.NewServer(
		ctx,
		makeHelloWorldEndpoint(svc),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/hello_service", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
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

type HelloWorldService interface {
	helloService(string) (string, error)
}

type helloWorldService struct {}

func (helloWorldService) helloService(name string) (string, error) {
	if name == "" {
		return "", ErrEmpty
	}
	return "Hello " + name, nil
}

var ErrEmpty = errors.New("Empty string")

type helloWorldRequest struct {
	Name string `json:"Name"`
}

type helloWorldResponse struct {
	ResponseText   string `json:"response"`
	Err string `json:"err,omitempty"`
}

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
package util

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
	http2 "github.com/go-kit/kit/transport/http"
)

func MakeDecoder(data CreateNewDatatype) http2.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {

		body := data()

		if err := json.NewDecoder(r.Body).Decode(body); err != nil {
			return nil, err
		}

		request := Request{r,body}
		return request, nil
	}
}

type CreateNewDatatype func() interface{}

type Request struct {
	*http.Request
	Data interface{}
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}





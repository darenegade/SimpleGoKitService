package util

import (
	"encoding/json"
	"net/http"
	"strconv"

	"encoding/xml"
	http2 "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

var ID = "id"
var IDPATH = "/{" + ID + "}"

func MakeDecoder(data CreateNewDatatype) http2.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {

		request := Request{r, nil, 0}

		if r.Method != http.MethodGet {

			body := data()

			if r.Header.Get("Content-Type") == "application/json" {

				if err := json.NewDecoder(r.Body).Decode(body); err != nil {
					return nil, err
				}
			} else if r.Header.Get("Content-Type") == "application/xml" {

				if err := xml.NewDecoder(r.Body).Decode(body); err != nil {
					return nil, err
				}
			} else {
				return nil, ErrUnsupportedMediaType
			}

			request.Data = body
		}

		return request, nil
	}
}

func MakePathIDDecoder(data CreateNewDatatype) http2.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		vars := mux.Vars(r)
		idstring, ok := vars[ID]

		if !ok {
			return nil, ErrBadRoute
		}

		parsedID, err := strconv.ParseUint(idstring, 10, 32)

		if err != nil {
			return nil, ErrBadRoute
		}

		id := uint(parsedID)

		if r.Method != http.MethodGet && r.Method != http.MethodDelete {

			body := data()

			if r.Header.Get("Content-Type") == "application/json" {

				if err := json.NewDecoder(r.Body).Decode(body); err != nil {
					return nil, err
				}
			} else if r.Header.Get("Content-Type") == "application/xml" {

				if err := xml.NewDecoder(r.Body).Decode(body); err != nil {
					return nil, err
				}
			} else {
				return nil, ErrUnsupportedMediaType
			}

			return Request{r, body, id}, nil

		} else {
			return Request{r, nil, id}, nil
		}
	}
}

type CreateNewDatatype func() interface{}

type Request struct {
	*http.Request
	Data interface{}
	ID   uint
}

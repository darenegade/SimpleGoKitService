package employee

import (
	"github.com/go-kit/kit/endpoint"
	http2 "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"net/http"

	"github.com/darenegade/SimpleGoKitService/database"
	"github.com/darenegade/SimpleGoKitService/util"
)

var svc = EmployeeService{}

var PATH = "/employees"

func MakeEmployeesEndpoint() (string, endpoint.Endpoint) {

	return PATH, func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(util.Request)

		if req.Method == http.MethodPost {
			return svc.create(*(req.Data.(*database.Employee)))
		} else if req.Method == http.MethodGet {
			return svc.findAll()
		} else {
			return nil, util.ErrWrongMethod
		}
	}
}

func MakeEmployeeEndpoint() (string, endpoint.Endpoint) {

	return PATH + util.IDPATH, func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(util.Request)

		if req.Method == http.MethodGet {
			return svc.findOne(req.ID)
		} else if req.Method == http.MethodDelete {
			return "", svc.delete(req.ID)
		} else if req.Method == http.MethodPut {
			return svc.update(*(req.Data.(*database.Employee)), req.ID)
		} else {
			return nil, util.ErrWrongMethod
		}
	}
}

func MakeEmployeesDecoder() http2.DecodeRequestFunc {
	return util.MakeDecoder(func() interface{} {
		return new(database.Employee)
	})
}

func MakeEmployeeDecoder() http2.DecodeRequestFunc {
	return util.MakePathIDDecoder(func() interface{} {
		return new(database.Employee)
	})
}

package main

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func Logging (logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {

			data,err := next(ctx, request)

			if err != nil {
				defer func(begin time.Time) {
					logger.Log(
						"Time", time.Now(),
						"err", err,
						"took", time.Since(begin),
					)
				}(time.Now())
			}

			return data,err
		}
	}
}

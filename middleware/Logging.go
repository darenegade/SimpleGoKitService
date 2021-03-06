package middleware

import (
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
)

func Logging(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			data, err := next(ctx, request)

			if err != nil {
				defer func(begin time.Time) {
					logger.Log(
						"Time", time.Now(),
						"err", err,
						"took", time.Since(begin),
					)
				}(time.Now())
			}

			return data, err
		}
	}
}

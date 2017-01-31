package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   HelloWorldService
}

func (mw loggingMiddleware) helloService(name string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Hello_World",
			"input", name,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.helloService(name)
	return
}

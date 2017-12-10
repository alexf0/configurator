package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) LoadParams(tp, data string) (c *Params, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "loadParams",
			"type", tp,
			"data", data,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.LoadParams(tp, data)
}

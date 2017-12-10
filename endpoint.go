package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type loadParamsRequest struct {
	Type string
	Data string
}

type loadErrorResponse struct {
	Err    error   `json:"error,omitempty"`
}

func (r loadErrorResponse) error() error {
	return r.Err
}

func makeLoadParamsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loadParamsRequest)
		params, err := s.LoadParams(req.Type, req.Data)

		if err != nil {
			return loadErrorResponse{err}, nil
		}

		return params, nil
	}
}


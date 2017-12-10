package main

import "errors"

var ErrInvalidArgument = errors.New("invalid argument")


type Service interface {
	LoadParams(tp, data string) (*Params, error)
}

type service struct {
	configs Repository
}

func (s *service) LoadParams(tp, data string) (*Params, error) {
	if tp == "" || data == "" {
		return nil, ErrInvalidArgument
	}

	c, err := s.configs.GetConfig(tp, data)

	if err != nil {
		return nil, err
	}

	return &c.Params, err
}

func NewService(configs Repository) Service {
	return &service{configs: configs}
}



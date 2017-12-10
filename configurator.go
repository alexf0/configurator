package main

import "errors"

var ErrUnknown = errors.New("unknown config")

type Params map[string] interface{}

type Config struct {
	Type string
	Data string
	Params Params `json:"params"`
}

type Repository interface {
	GetConfig(tp, data string) (*Config, error)
	Store(config *Config) error
}
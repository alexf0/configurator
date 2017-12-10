package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var testService Service

type testRepository struct {
	config *Config
}

func (r *testRepository) GetConfig(tp, data string) (*Config, error) {
	if r.config.Type == tp && r.config.Data == data {
		return r.config, nil
	}

	return nil, ErrUnknown
}

func (r *testRepository) Store(config *Config) error {
	r.config  = config

	return nil
}

func init() {
	rep := &testRepository{}
	testStore(rep)
	testService = NewService(rep)
}

func TestServiceEmptyParams(t *testing.T) {
	tp, data := "", ""

	_, err := testService.LoadParams(tp, data)

	assert.Equal(t, err, ErrInvalidArgument, "errors should be equal")
}

func TestServiceBadParams(t *testing.T) {
	tp, data := "Type1", "Data1"

	params, _ := testService.LoadParams(tp, data)

	assert.Nil(t, params, "config should be equal nil")
}

func TestServiceSuccessParams(t *testing.T) {
	tp, data := "Type", "Data"

	from, _ := testService.LoadParams(tp, data)

	to := map[string]interface{} {
		"param1": "param1",
		"param2":  "param2",
	}

	assert.Equal(t, *from, Params(to), "should be equal")
}

func testStore(r Repository) *Config {
	params := map[string]interface{} {
		"param1": "param1",
		"param2":  "param2",
	}

	c := &Config{
		Type: "Type",
		Data: "Data",
		Params: Params(params),
	}

	r.Store(c)

	return c
}

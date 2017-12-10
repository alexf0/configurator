package main

import (
	"testing"
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/stretchr/testify/assert"
)

var repPgTest Repository

func init() {
	conf, err := yaml.ReadFile("conf.yml")

	if err != nil {
		panic(err)
	}

	repPgTest, err = NewConfigRepository(conf)

	if err != nil {
		panic(err)
	}
}

func TestFoundRecord(t *testing.T) {
	tp, data := "FoundType", "FoundData"

	params := map[string]interface{} {
		"param1": "param1",
		"param2":  "param2",
	}

	storeConfig := &Config{
		Type: tp,
		Data: data,
		Params: Params(params),
	}

	err := repPgTest.Store(storeConfig)

	if err != nil {
		t.Error(err)
	}

	config, _ := repPgTest.GetConfig(tp, data)

	assert.Equal(t, config, storeConfig, "should be equal")
}

func TestNotFoundRecord(t *testing.T) {
	tp1, tp2 := "FoundType", "NotFoundType"
	data1, data2 := "FoundData", "NotFoundData"

	params := map[string]interface{} {
		"param1": "param1",
		"param2":  "param2",
	}

	storeConfig := &Config{
		Type: tp1,
		Data: data1,
		Params: Params(params),
	}

	repPgTest.Store(storeConfig)

	config, err := repPgTest.GetConfig(tp1, data2)

	assert.Nil(t, config)
	assert.Equal(t, err, ErrUnknown, "errors should be equal")

	config, err = repPgTest.GetConfig(tp2, data1)

	assert.Nil(t, config)
	assert.Equal(t, err, ErrUnknown, "errors should be equal")

	config, err = repPgTest.GetConfig(tp2, data2)

	assert.Nil(t, config)
	assert.Equal(t, err, ErrUnknown, "errors should be equal")
}


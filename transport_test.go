package main

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"bytes"
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestServerBadRequest(t *testing.T) {
	logger := log.NewLogfmtLogger(&bytes.Buffer{})
	rep := &testRepository{}
	sc := NewService(rep)
	server := httptest.NewServer(MakeHandler(sc, logger))
	url := server.URL+"/api/v1/params"

	defer server.Close()

	var json1 = []byte(`{}`)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(json1))

	if want, have := http.StatusBadRequest, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}

	var json2 = []byte(`{"Type":"","Data":""}`)
	resp, _ = http.Post(url, "application/json", bytes.NewBuffer(json2))

	if want, have := http.StatusBadRequest, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}

	var json3 = []byte(`{"Type":"Type","Data":""}`)
	resp, _ = http.Post(url, "application/json", bytes.NewBuffer(json3))

	if want, have := http.StatusBadRequest, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}

	var json4 = []byte(`{"Type":"","Data":"Data"}`)
	resp, _ = http.Post(url, "application/json", bytes.NewBuffer(json4))

	if want, have := http.StatusBadRequest, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerNotFound(t *testing.T) {
	logger := log.NewLogfmtLogger(&bytes.Buffer{})
	rep := &testRepository{}
	testStore(rep)
	sc := NewService(rep)
	server := httptest.NewServer(MakeHandler(sc, logger))
	url := server.URL+"/api/v1/params"

	defer server.Close()

	var json1 = []byte(`{"Type":"TypeNot","Data":"DataNot"}`)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(json1))

	if want, have := http.StatusNotFound, resp.StatusCode; want != have {
		t.Errorf("want %d, have %d", want, have)
	}
}

func TestServerFound(t *testing.T) {
	logger := log.NewLogfmtLogger(&bytes.Buffer{})
	rep := &testRepository{}
	c := testStore(rep)
	sc := NewService(rep)
	server := httptest.NewServer(MakeHandler(sc, logger))
	url := server.URL+"/api/v1/params"

	defer server.Close()

	var json1 = []byte(`{"Type":"Type","Data":"Data"}`)
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(json1))

	if assert.Equal(t, resp.StatusCode, http.StatusOK) {
		var params Params
		json.NewDecoder(resp.Body).Decode(&params)
		assert.Equal(t, params, c.Params)
	}
}

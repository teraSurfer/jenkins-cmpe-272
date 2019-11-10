package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", Status).Methods("GET")
	return r
}

func TestStatus(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	Router().ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code, "OK response is expected")

}

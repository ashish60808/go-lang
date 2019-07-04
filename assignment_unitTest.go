package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	router.HandleFunc("/configs", GetConfigs).Methods("GET")
	router.HandleFunc("/configs/{name}", GetConfig).Methods("GET")
	router.HandleFunc("/configs", CreateConfig).Methods("POST")
	router.HandleFunc("/configs/{name}", DeleteConfig).Methods("DELETE")
	router.HandleFunc("/search", SearchConfig).Methods("GET")
	return router
}

func TestHealthcheckEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/healthcheck", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, "Application is up and running\n", response.Body.String(), "Correct response is found")
	// fmt.Println(response.Body)
}

func TestListConfigsEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/configs", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

func TestGetConfigEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/configs/config1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

func TestCreateConfigEndpoint(t *testing.T) {
	payload := []byte(`{"name":"config5","data":{"key":"val5", "id":5}}`)
	request, _ := http.NewRequest("POST", "/configs", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

func TestDeleteConfigEndpoint(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/configs/config1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

func TestSearchConfigEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/search?name=config2&data.id=2", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

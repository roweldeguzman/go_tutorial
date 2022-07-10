package main

import (
	"api/server"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app server.App

func TestMain(m *testing.M) {
	app = server.App{}
	app.Initialize()
	code := m.Run()
	os.Exit(code)
}

func executeRequest(request *http.Request, handler http.HandlerFunc) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	return response
}

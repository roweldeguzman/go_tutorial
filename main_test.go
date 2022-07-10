package main

import (
	"api/controllers"
	"api/controllers/users"
	"api/server"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestLogin(t *testing.T) {

	var jsonStr = []byte(`{"email":"rowel1@gmail.com", "password": "admin"}`)
	request, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(controllers.Login)

	response := executeRequest(request, handler)

	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)

	if err != nil {
		t.Errorf("Unexpected response body found")
	}
	if m["statusCode"] != float64(200) {
		t.Errorf("Wrong status code. expecting 200, receive %v", m["statusCode"])
	}
}

func TestGetUsers(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/users/get", nil)
	query := request.URL.Query()
	query.Add("rows", "1")
	query.Add("page", "2")
	request.URL.RawQuery = query.Encode()

	handler := http.HandlerFunc(users.Get)
	response := executeRequest(request, handler)

	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)

	if err != nil {
		t.Errorf("Unexpected response body found")
	}

	devMessage := m["devMessage"].([]interface{})
	paginate := m["paginate"].(map[string]interface{})

	if m["statusCode"] != float64(200) {
		t.Errorf("Wrong status code. expecting 200, got %v", m["statusCode"])
	}

	if len(devMessage) > 1 {
		t.Errorf("wrong query result. expecting %v. got %v", 1, len(devMessage))
	}
	if paginate["currentPage"] != float64(2) {
		t.Errorf("wrong current page. expecting %v. got %v", 2, paginate["currentPage"])
	}

	assert.Equal(t, float64(2), paginate["currentPage"])

}

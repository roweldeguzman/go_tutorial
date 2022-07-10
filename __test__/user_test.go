package main

import (
	"api/controllers/users"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	ID int64 `json:"id"`
}

func TestLogin(t *testing.T) {

	//The response recorder used to record HTTP responses
	response := httptest.NewRecorder()
	var jsonStr = []byte(`{"email":"rowel1@gmail.com", "passwords": "admins"}`)
	request, err := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatal("Creating 'GET /questions/1/SC' request failed!")
	}

	app.Router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", response.Code, " instead of ", http.StatusOK)
	}
}

func TestCreates(t *testing.T) {

	var jsonStr = []byte(`{"firstName":"Rowel", "lastName": "de Guzman", "email": "rowel143deguzman@gmail.com", "userStatus": "1"}`)
	request, _ := http.NewRequest("POST", "/v1/users/add", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users.Create(w, r, nil)
	})

	response := executeRequest(request, handler)

	var m map[string]interface{}
	err := json.Unmarshal(response.Body.Bytes(), &m)

	fmt.Println(m)
	if err != nil {
		t.Errorf("Unexpected response body found")
	}
	if m["statusCode"] != float64(200) {
		t.Errorf("Wrong status code. expecting 200, receive %v", m["statusCode"])
	}
}

func TestDelete(t *testing.T) {}

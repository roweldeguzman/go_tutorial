package main

import (
	controller "api/controllers"
	user "api/controllers/users"
	"api/repository"
	"api/service"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userRepository = repository.NewUserRepository()
	userService    = service.NewUserService(userRepository)
	UserController = user.NewUserController(userService)
	authController = controller.NewAuthController(userService)
)

func tearDown(t *testing.T, id string) {
	var jsonStr = []byte(`{"ids": [` + id + `]}`)
	request, _ := http.NewRequest("DELETE", "/v1/users/delete", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserController.Delete(w, r)
	})

	response := executeRequest(request, handler)

	var m map[string]any
	err := json.Unmarshal(response.Body.Bytes(), &m)

	if err != nil {
		t.Errorf("Failed to delete user with ID" + id)
	}
}

func TestLogin(t *testing.T) {

	var jsonStr = []byte(`{"email":"rowel@gmail.com", "password": "admin"}`)
	request, _ := http.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(authController.Login)

	response := executeRequest(request, handler)

	var m map[string]any
	err := json.Unmarshal(response.Body.Bytes(), &m)

	if err != nil {
		t.Errorf("Unexpected response body found")
	}
	if m["statusCode"] != float64(200) {
		t.Errorf("Wrong status code. expecting 200, receive %v", m["statusCode"])
	}
}

func TestCreates(t *testing.T) {

	var jsonStr = []byte(`{"firstName":"Rowel", "lastName": "de Guzman", "email": "` + faker.Email() + `", "password": "password1",  "userStatus": "1"}`)
	request, _ := http.NewRequest("POST", "/v1/users/add", bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		UserController.Create(w, r)
	})

	response := executeRequest(request, handler)

	var m map[string]any
	err := json.Unmarshal(response.Body.Bytes(), &m)

	if err != nil {
		t.Errorf("Unexpected response body found")
	}
	if m["statusCode"] != float64(200) {
		t.Errorf("Wrong status code. expecting 200, receive %v", m["statusCode"])
	}

	devMessage := m["devMessage"].(map[string]any)

	fmt.Println(devMessage)
	id, _ := devMessage["id"].(float64)
	tearDown(t, strconv.Itoa(int(id)))

}

func TestGetUsers(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/users/get", nil)
	query := request.URL.Query()
	query.Add("rows", "1")
	query.Add("page", "2")
	request.URL.RawQuery = query.Encode()

	handler := http.HandlerFunc(UserController.Get)
	response := executeRequest(request, handler)

	var m map[string]any
	err := json.Unmarshal(response.Body.Bytes(), &m)

	if err != nil {
		t.Errorf("Unexpected response body found")
	}

	devMessage := m["devMessage"].([]any)
	paginate := m["paginate"].(map[string]any)

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

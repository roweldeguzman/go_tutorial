package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:9999/v1" // Replace with your actual server URL

func TestE2E_Login(t *testing.T) {
	// Prepare login payload
	var jsonStr = []byte(`{
        "email":"rowel.deguzman@roweldev.com",
        "password": "admin"
    }`)

	// Send POST request to login endpoint
	response, err := http.Post(baseURL+"/auth/login", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()

	// Parse response
	var m map[string]any
	err = json.NewDecoder(response.Body).Decode(&m)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Assertions
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(200), m["statusCode"])
}

func TestE2E_CreateUser(t *testing.T) {
	// Generate fake user data
	firstName := faker.FirstName()
	lastName := faker.LastName()
	email := faker.Email()

	// Prepare user creation payload
	var jsonStr = []byte(`{
        "firstName":"` + firstName + `",
        "lastName": "` + lastName + `",
        "email": "` + email + `",
        "password": "password1",
        "userStatus": "1"
    }`)

	// Send POST request to create user endpoint
	response, err := http.Post(baseURL+"/users/add", "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()

	// Parse response
	var m map[string]any
	err = json.NewDecoder(response.Body).Decode(&m)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Assertions
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(200), m["statusCode"])

	devMessage := m["devMessage"].(map[string]any)
	assert.Equal(t, firstName, devMessage["firstName"])
	assert.Equal(t, lastName, devMessage["lastName"])
	assert.Equal(t, email, devMessage["email"])

	// Clean up: Delete the created user
	id, _ := devMessage["id"].(float64)
	tearDownE2E(t, strconv.Itoa(int(id)))
}

func TestE2E_GetUsers(t *testing.T) {
	// Prepare GET request with query parameters
	req, err := http.NewRequest("GET", baseURL+"/users/get", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	query := req.URL.Query()
	query.Add("rows", "1")
	query.Add("page", "2")
	req.URL.RawQuery = query.Encode()

	// Send GET request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()

	// Parse response
	var m map[string]any
	err = json.NewDecoder(response.Body).Decode(&m)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Assertions
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, float64(200), m["statusCode"])

	devMessage := m["devMessage"].([]any)
	paginate := m["paginate"].(map[string]any)

	assert.Equal(t, 1, len(devMessage))
	assert.Equal(t, float64(2), paginate["currentPage"])
}

func tearDownE2E(t *testing.T, id string) {
	// Prepare DELETE request payload
	var jsonStr = []byte(`{"ids": [` + id + `]}`)
	req, err := http.NewRequest("DELETE", baseURL+"/users/delete", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send DELETE request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer response.Body.Close()

	// Parse response
	var m map[string]any
	err = json.NewDecoder(response.Body).Decode(&m)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Assertions
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

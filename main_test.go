package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBasicAssertion(t *testing.T) {
	// Simple test that asserts 1==1
	if 1 != 1 {
		t.Error("Expected 1 to equal 1")
	}
}

func TestHelloWorldEndpoint(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	router := getRouter()

	// Create a test request
	req, _ := http.NewRequest("GET", "/hello_world", nil)
	resp := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(resp, req)

	// Check the status code
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := map[string]string{"hello": "world"}
	var actual map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &actual)
	if err != nil {
		t.Fatal(err)
	}

	if actual["hello"] != expected["hello"] {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}

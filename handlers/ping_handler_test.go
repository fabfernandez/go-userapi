package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	handler := NewPingHandler()
	
	// Create a request to pass to our handler
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	// Call the handler
	handler.Ping(w, req)

	// Check the status code
	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	// Check the response body
	var response map[string]string
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Could not decode response body: %v", err)
	}

	// Check if the message is correct
	expectedMessage := "pong"
	if response["message"] != expectedMessage {
		t.Errorf("handler returned unexpected message: got %v want %v", response["message"], expectedMessage)
	}
} 
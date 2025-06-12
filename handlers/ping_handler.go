package handlers

import (
	"encoding/json"
	"net/http"
)

// PingHandler handles health check requests
type PingHandler struct{}

// NewPingHandler creates a new ping handler
func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

// Ping handles the ping request
// @Summary Health check endpoint
// @Description Returns a simple pong message to verify the API is running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "pong",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
} 
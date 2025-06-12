package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"userapi/models"
	"userapi/repository"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	if err := user.Validate(); err != nil {
		log.Printf("Validation error for user: %v", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.repo.Create(r.Context(), &user); err != nil {
		log.Printf("Error creating user: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	log.Printf("Successfully created user with ID: %d", user.ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// @Summary Get a user by ID
// @Description Get a user by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		log.Printf("Error retrieving user with ID %d: %v", id, err)
		respondWithError(w, http.StatusInternalServerError, "Error retrieving user")
		return
	}
	if user == nil {
		log.Printf("User not found with ID: %d", id)
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	log.Printf("Successfully retrieved user with ID: %d", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// @Summary Update a user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [put]
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding request body: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user.ID = id
	if err := user.Validate(); err != nil {
		log.Printf("Validation error for user update: %v", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.repo.Update(r.Context(), &user); err != nil {
		log.Printf("Error updating user with ID %d: %v", id, err)
		if err.Error() == fmt.Sprintf("user not found with ID: %d", id) {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Error updating user")
		return
	}

	log.Printf("Successfully updated user with ID: %d", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		log.Printf("Error deleting user with ID %d: %v", id, err)
		if err.Error() == fmt.Sprintf("user not found with ID: %d", id) {
			respondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Error deleting user")
		return
	}

	log.Printf("Successfully deleted user with ID: %d", id)
	w.WriteHeader(http.StatusNoContent)
}

// @Summary List all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.List(r.Context())
	if err != nil {
		log.Printf("Error retrieving users list: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error retrieving users")
		return
	}

	log.Printf("Successfully retrieved %d users", len(users))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	log.Printf("Responding with error: %s (status code: %d)", message, code)
	respondWithJSON(w, code, ErrorResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
} 
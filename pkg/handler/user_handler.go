package handler

import (
	"activity-tracker/pkg/model"
	repository "activity-tracker/pkg/respository"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// UserHandler handles HTTP requests related to users.
type UserHandler struct {
	userRepo *repository.Repository
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(userRepo *repository.Repository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// RegisterRoutes registers the user routes.
func (h *UserHandler) RegisterRoutes(router chi.Router) {
	router.Post("/users", h.CreateUser)
	router.Get("/users/{userID}", h.GetUser)
	router.Put("/users/{userID}", h.UpdateUser)
	router.Delete("/users/{userID}", h.DeleteUser)
}

// CreateUser handles the creation of a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := h.userRepo.CreateUser(&user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := map[string]int64{"user_id": userID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUser handles retrieving a user by ID.
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUser(userID)
	if errors.Is(err, repository.ErrUserNotFound) {
		log.Printf("User not found: %v", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error retrieving user: %v", err)
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser handles updating a user by ID.
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user.ID = userID

	if err := h.userRepo.UpdateUser(&user); err != nil {
		log.Printf("Error updating user: %v", err)
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		log.Printf("Error parsing user ID: %v", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.userRepo.DeleteUser(userID); err != nil {
		log.Printf("Error deleting user: %v", err)
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

package handler

import (
	"activity-tracker/pkg/model"
	repository "activity-tracker/pkg/respository"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// UserActivityHandler handles HTTP requests related to user activities.
type UserActivityHandler struct {
	userActivityRepo *repository.Repository
}

// NewUserActivityHandler creates a new UserActivityHandler instance.
func NewUserActivityHandler(userActivityRepo *repository.Repository) *UserActivityHandler {
	return &UserActivityHandler{userActivityRepo: userActivityRepo}
}

// RegisterRoutes registers the user activity routes.
func (h *UserActivityHandler) RegisterRoutes(router chi.Router) {
	router.Post("/user-activities", h.CreateUserActivity)
	router.Get("/user-activities/{userActivityID}", h.GetUserActivity)
	router.Put("/user-activities/{userActivityID}", h.UpdateUserActivity)
	router.Delete("/user-activities/{userActivityID}", h.DeleteUserActivity)
}

// CreateUserActivity handles the creation of a new user activity.
func (h *UserActivityHandler) CreateUserActivity(w http.ResponseWriter, r *http.Request) {
	var userActivity model.UserActivity
	if err := json.NewDecoder(r.Body).Decode(&userActivity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userActivityID, err := h.userActivityRepo.CreateUserActivity(&userActivity)
	if err != nil {
		http.Error(w, "Failed to create user activity", http.StatusInternalServerError)
		return
	}

	response := map[string]int64{"user_activity_id": userActivityID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUserActivity handles retrieving a user activity by ID.
func (h *UserActivityHandler) GetUserActivity(w http.ResponseWriter, r *http.Request) {
	userActivityID, err := strconv.ParseInt(chi.URLParam(r, "userActivityID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user activity ID", http.StatusBadRequest)
		return
	}

	userActivity, err := h.userActivityRepo.GetUserActivity(userActivityID)
	if errors.Is(err, repository.ErrUserActivityNotFound) {
		http.Error(w, "User activity not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to retrieve user activity", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userActivity)
}

// UpdateUserActivity handles updating a user activity by ID.
func (h *UserActivityHandler) UpdateUserActivity(w http.ResponseWriter, r *http.Request) {
	userActivityID, err := strconv.ParseInt(chi.URLParam(r, "userActivityID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user activity ID", http.StatusBadRequest)
		return
	}

	var userActivity model.UserActivity
	if err := json.NewDecoder(r.Body).Decode(&userActivity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userActivity.ID = userActivityID

	if err := h.userActivityRepo.UpdateUserActivity(&userActivity); err != nil {
		http.Error(w, "Failed to update user activity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUserActivity handles deleting a user activity by ID.
func (h *UserActivityHandler) DeleteUserActivity(w http.ResponseWriter, r *http.Request) {
	userActivityID, err := strconv.ParseInt(chi.URLParam(r, "userActivityID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid user activity ID", http.StatusBadRequest)
		return
	}

	if err := h.userActivityRepo.DeleteUserActivity(userActivityID); err != nil {
		http.Error(w, "Failed to delete user activity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

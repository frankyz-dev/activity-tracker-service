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

// ActivityHandler handles HTTP requests related to activities.
type ActivityHandler struct {
	activityRepo *repository.Repository
}

// NewActivityHandler creates a new ActivityHandler instance.
func NewActivityHandler(activityRepo *repository.Repository) *ActivityHandler {
	return &ActivityHandler{activityRepo: activityRepo}
}

// RegisterRoutes registers the activity routes.
func (h *ActivityHandler) RegisterRoutes(router chi.Router) {
	router.Post("/activities", h.CreateActivity)
	router.Get("/activities/{activityID}", h.GetActivity)
	router.Put("/activities/{activityID}", h.UpdateActivity)
	router.Delete("/activities/{activityID}", h.DeleteActivity)
}

// CreateActivity handles the creation of a new activity.
func (h *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	var activity model.Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	activityID, err := h.activityRepo.CreateActivity(&activity)
	if err != nil {
		http.Error(w, "Failed to create activity", http.StatusInternalServerError)
		return
	}

	response := map[string]int64{"activity_id": activityID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetActivity handles retrieving an activity by ID.
func (h *ActivityHandler) GetActivity(w http.ResponseWriter, r *http.Request) {
	activityID, err := strconv.ParseInt(chi.URLParam(r, "activityID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid activity ID", http.StatusBadRequest)
		return
	}

	activity, err := h.activityRepo.GetActivity(activityID)
	if errors.Is(err, repository.ErrActivityNotFound) {
		http.Error(w, "Activity not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Failed to retrieve activity", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
}

// UpdateActivity handles updating an activity by ID.
func (h *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	activityID, err := strconv.ParseInt(chi.URLParam(r, "activityID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid activity ID", http.StatusBadRequest)
		return
	}

	var activity model.Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	activity.ID = activityID

	if err := h.activityRepo.UpdateActivity(&activity); err != nil {
		http.Error(w, "Failed to update activity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteActivity handles deleting an activity by ID.
func (h *ActivityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	activityID, err := strconv.ParseInt(chi.URLParam(r, "activityID"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid activity ID", http.StatusBadRequest)
		return
	}

	if err := h.activityRepo.DeleteActivity(activityID); err != nil {
		http.Error(w, "Failed to delete activity", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

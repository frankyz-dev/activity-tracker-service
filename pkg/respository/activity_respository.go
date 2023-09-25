package repository

import (
	"activity-tracker/pkg/model"
	"database/sql"
	"errors"
	"fmt"
)

// ErrActivityNotFound is returned when the activity is not found in the database.
var ErrActivityNotFound = errors.New("activity not found")

// CreateActivity creates a new activity in the database.
func (r *Repository) CreateActivity(activity *model.Activity) (int64, error) {
	var id int64
	query := `INSERT INTO activities (name) VALUES ($1) RETURNING activity_id`
	err := r.db.QueryRow(query, activity.Name).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not create activity: %w", err)
	}
	return id, nil
}

// GetActivity retrieves an activity by ID from the database.
func (r *Repository) GetActivity(activityID int64) (*model.Activity, error) {
	activity := &model.Activity{}
	query := `SELECT activity_id, name FROM activities WHERE activity_id = $1`
	err := r.db.QueryRow(query, activityID).Scan(&activity.ID, &activity.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrActivityNotFound
		}
		return nil, fmt.Errorf("could not get activity: %w", err)
	}
	return activity, nil
}

// UpdateActivity updates an existing activity in the database.
func (r *Repository) UpdateActivity(activity *model.Activity) error {
	query := `UPDATE activities SET name = $1 WHERE activity_id = $2`
	_, err := r.db.Exec(query, activity.Name, activity.ID)
	if err != nil {
		return fmt.Errorf("could not update activity: %w", err)
	}
	return nil
}

// DeleteActivity deletes an activity by ID from the database.
func (r *Repository) DeleteActivity(activityID int64) error {
	query := `DELETE FROM activities WHERE activity_id = $1`
	_, err := r.db.Exec(query, activityID)
	if err != nil {
		return fmt.Errorf("could not delete activity: %w", err)
	}
	return nil
}

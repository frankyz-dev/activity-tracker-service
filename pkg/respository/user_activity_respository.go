package repository

import (
	"activity-tracker/pkg/model"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// ErrUserActivityNotFound is returned when the user activity is not found in the database.
var ErrUserActivityNotFound = errors.New("user activity not found")

// CreateUserActivity creates a new user activity in the database.
func (r *Repository) CreateUserActivity(userActivity *model.UserActivity) (int64, error) {
	var id int64
	additionalAttributes, err := json.Marshal(userActivity.AdditionalAttributes)
	if err != nil {
		return 0, fmt.Errorf("could not marshal additional attributes: %w", err)
	}

	query := `INSERT INTO user_activities (user_id, activity_id, start_time, end_time, duration, mood, additional_attributes, recorded_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err = r.db.QueryRow(query, userActivity.UserID, userActivity.ActivityID, userActivity.StartTime, userActivity.EndTime,
		userActivity.Duration, userActivity.Mood, additionalAttributes, time.Now()).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not create user activity: %w", err)
	}
	return id, nil
}

// GetUserActivity retrieves a user activity by ID from the database.
func (r *Repository) GetUserActivity(userActivityID int64) (*model.UserActivity, error) {
	userActivity := &model.UserActivity{}
	var additionalAttributes []byte
	query := `SELECT id, user_id, activity_id, start_time, end_time, duration, mood, additional_attributes, recorded_at
			  FROM user_activities WHERE id = $1`
	err := r.db.QueryRow(query, userActivityID).Scan(&userActivity.ID, &userActivity.UserID, &userActivity.ActivityID,
		&userActivity.StartTime, &userActivity.EndTime, &userActivity.Duration, &userActivity.Mood, &additionalAttributes, &userActivity.RecordedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserActivityNotFound
		}
		return nil, fmt.Errorf("could not get user activity: %w", err)
	}

	err = json.Unmarshal(additionalAttributes, &userActivity.AdditionalAttributes)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal additional attributes: %w", err)
	}
	return userActivity, nil
}

// UpdateUserActivity updates an existing user activity in the database.
func (r *Repository) UpdateUserActivity(userActivity *model.UserActivity) error {
	additionalAttributes, err := json.Marshal(userActivity.AdditionalAttributes)
	if err != nil {
		return fmt.Errorf("could not marshal additional attributes: %w", err)
	}

	query := `UPDATE user_activities SET start_time = $1, end_time = $2, duration = $3, mood = $4, additional_attributes = $5
			  WHERE id = $6`
	_, err = r.db.Exec(query, userActivity.StartTime, userActivity.EndTime, userActivity.Duration, userActivity.Mood, additionalAttributes, userActivity.ID)
	if err != nil {
		return fmt.Errorf("could not update user activity: %w", err)
	}
	return nil
}

// DeleteUserActivity deletes a user activity by ID from the database.
func (r *Repository) DeleteUserActivity(userActivityID int64) error {
	query := `DELETE FROM user_activities WHERE id = $1`
	_, err := r.db.Exec(query, userActivityID)
	if err != nil {
		return fmt.Errorf("could not delete user activity: %w", err)
	}
	return nil
}

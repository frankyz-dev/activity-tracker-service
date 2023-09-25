package repository

import (
	"activity-tracker/pkg/model"
	"database/sql"
	"fmt"
)

// CreateUser creates a new user in the database.
func (r *Repository) CreateUser(user *model.User) (int64, error) {
	var id int64
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not create user: %w", err)
	}
	return id, nil
}

// GetUser retrieves a user by ID from the database.
func (r *Repository) GetUser(userID int64) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, username, password, created_at FROM users WHERE id = $1`
	err := r.db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("could not get user: %w", err)
	}
	return user, nil
}

// UpdateUser updates an existing user in the database.
func (r *Repository) UpdateUser(user *model.User) error {
	query := `UPDATE users SET username = $1, password = $2 WHERE id = $3`
	_, err := r.db.Exec(query, user.Username, user.Password, user.ID)
	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}
	return nil
}

// DeleteUser deletes a user by ID from the database.
func (r *Repository) DeleteUser(userID int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}
	return nil
}

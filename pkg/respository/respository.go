package repository

import (
	"database/sql"
	"errors"
)

// Repository provides methods to interact with the database.
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance.
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ErrUserNotFound is returned when the user is not found in the database.
var ErrUserNotFound = errors.New("user not found")

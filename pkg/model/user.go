package model

import "time"

// User represents the user data model.
type User struct {
	ID        int64     `db:"user_id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"` // Store hashed password, not plain text
	CreatedAt time.Time `db:"created_at"`
}

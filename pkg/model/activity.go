package model

// Activity represents the activity data model.
type Activity struct {
	ID   int64  `db:"activity_id"`
	Name string `db:"name"`
	// Add other fields as needed, e.g., description, category, etc.
}

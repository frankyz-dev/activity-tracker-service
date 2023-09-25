package model

import (
	"encoding/json"
	"time"
)

// AdditionalAttributes represents the structure of the additional attributes in JSONB format.
type AdditionalAttributes struct {
	KneeFeeling string `json:"knee_feeling,omitempty"`
	// Add other fields as needed
}

// UserActivity represents a user's activity record.
type UserActivity struct {
	ID                   int64
	UserID               int64 `db:"user_id"`
	ActivityID           int64 `db:"activity_id"`
	StartTime            time.Time
	EndTime              time.Time
	Duration             time.Duration
	Mood                 int
	AdditionalAttributes AdditionalAttributes
	RecordedAt           time.Time `db:"recorded_at"`
}

// MarshalAdditionalAttributes marshals AdditionalAttributes to JSONB format.
func (ua *UserActivity) MarshalAdditionalAttributes() ([]byte, error) {
	return json.Marshal(ua.AdditionalAttributes)
}

// UnmarshalAdditionalAttributes unmarshals AdditionalAttributes from JSONB format.
func (ua *UserActivity) UnmarshalAdditionalAttributes(data []byte) error {
	return json.Unmarshal(data, &ua.AdditionalAttributes)
}

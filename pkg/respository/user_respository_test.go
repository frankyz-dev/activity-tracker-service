package repository

import (
	"activity-tracker/pkg/model"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/stretchr/testify/assert"
)

// TODO: Change this to use in memory sqlite database
const (
	host     = "192.168.1.32"
	port     = 5432
	user     = "frankyz"
	password = "5LQimPMPxjMdB1wlJJBX"
	dbname   = "mudkip"
)

func TestCreateUser(t *testing.T) {
	// Setup test database connection
	psqlInfo := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	psqlInfo = fmt.Sprintf(psqlInfo, host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	repo := NewRepository(db)

	// Define test user
	testUser := &model.User{
		Username: "testUsername",
		Password: "testPassword", // In real scenario, store hashed password
	}

	// Execute CreateUser function
	userID, err := repo.CreateUser(testUser)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Assert user ID is returned and is valid
	assert.NotZero(t, userID)

	// Retrieve the created user and compare
	retrievedUser, err := repo.GetUser(userID)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %v", err)
	}
	assert.Equal(t, testUser.Username, retrievedUser.Username)

	// Cleanup: Delete the test user from the database
	err = repo.DeleteUser(userID)
	if err != nil {
		t.Fatalf("Failed to delete test user: %v", err)
	}
}

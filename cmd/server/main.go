package main

import (
	"activity-tracker/pkg/config"
	"activity-tracker/pkg/handler"
	repository "activity-tracker/pkg/respository"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)

func main() {
	db := ConnectToDatabase()
	defer db.Close()

	// Initialize repositories
	repo := repository.NewRepository(db)

	// Initialize handlers
	userHandler := handler.NewUserHandler(repo)
	activityHandler := handler.NewActivityHandler(repo)
	userActivityHandler := handler.NewUserActivityHandler(repo)

	// Initialize router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Register routes
	userHandler.RegisterRoutes(router)
	activityHandler.RegisterRoutes(router)
	userActivityHandler.RegisterRoutes(router)

	// Start the HTTP server
	log.Println("Server is running on port 8089")
	log.Fatal(http.ListenAndServe(":8089", router))
}

func ConnectToDatabase() *sql.DB {
	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

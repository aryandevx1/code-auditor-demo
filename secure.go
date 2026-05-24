package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// UserRepository holds a reference to the database connection
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository initialises a UserRepository with the given DB connection
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserSafe fetches a user record by ID using a parameterised query.
// This function is safe from SQL injection.
func GetUserSafe(db *sql.DB, id string) {
	db.QueryRow("SELECT * FROM users WHERE id = $1", id)
}

// ServeSecureUserAPI starts an HTTP server that exposes the secure user lookup endpoint
func ServeSecureUserAPI(repo *UserRepository) {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}
		GetUserSafe(repo.db, id)
		fmt.Fprintf(w, "Fetched user: %s", id)
	})

	log.Println("Starting secure user API server on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

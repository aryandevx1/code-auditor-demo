package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// UserStore holds a reference to the database connection
type UserStore struct {
	db *sql.DB
}

// NewUserStore initialises a UserStore with the given DB connection
func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

// GetUser fetches a user record by ID from the database.
// WARNING: This function is intentionally vulnerable to SQL injection.
func GetUser(db *sql.DB, id string) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", id)
	db.QueryRow(query)
}

// ServeUserAPI starts an HTTP server that exposes the user lookup endpoint
func ServeUserAPI(store *UserStore) {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "missing id parameter", http.StatusBadRequest)
			return
		}
		GetUser(store.db, id)
		fmt.Fprintf(w, "Fetched user: %s", id)
	})

	log.Println("Starting user API server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func main() {
	db, err := sql.Open("postgres", "host=localhost user=admin dbname=appdb sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	store := NewUserStore(db)
	ServeUserAPI(store)
}

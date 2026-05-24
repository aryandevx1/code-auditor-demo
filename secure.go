// secure.go
package main

import "database/sql"

func GetUserSafe(db *sql.DB, id string) {
	db.QueryRow("SELECT * FROM users WHERE id = $1", id)
}

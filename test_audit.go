// test_audit.go
package main

import (
	"database/sql"
	"fmt"
)

func GetUserTest(db *sql.DB, id string) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", id)
	db.QueryRow(query)

}

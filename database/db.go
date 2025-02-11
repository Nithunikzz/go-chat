// database/db.go
package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(databaseURL string) {
	var err error
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func SaveMessage(clientID, message string) {
	_, err := DB.Exec("INSERT INTO messages (client_id, message) VALUES ($1, $2)", clientID, message)
	if err != nil {
		log.Println("Error saving message:", err)
	}
}

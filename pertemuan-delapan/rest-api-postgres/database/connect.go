package database

import (
	"database/sql"
	"log"
	"rest-api-postgres/config"

	_ "github.com/lib/pq"
)

func InitDB() {
	var err error

	DB, err = sql.Open("postgres", config.ConnectionString())
	if err != nil {
		log.Fatal("failed to connect database")
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connection Opened to Database")
}

func CloseDB() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Println("Error closing database:", err)
		} else {
			log.Println("Database connection closed.")
		}
	}
}

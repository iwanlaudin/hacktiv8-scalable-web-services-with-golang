package config

import "fmt"

const (
	// api config
	PORT = ":8080"

	// database config
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

func ConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

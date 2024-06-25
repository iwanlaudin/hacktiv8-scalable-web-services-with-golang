package auth

import "database/sql"

type Repository struct {
	db *sql.DB
}

func AuthService(db *sql.DB) Service {
	return &Repository{
		db,
	}
}

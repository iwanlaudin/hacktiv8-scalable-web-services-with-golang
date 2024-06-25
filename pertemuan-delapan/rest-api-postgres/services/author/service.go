package author

import "database/sql"

type repository struct {
	db *sql.DB
}

func NewAuthorService(db *sql.DB) Service {
	return &repository{
		db: db,
	}
}

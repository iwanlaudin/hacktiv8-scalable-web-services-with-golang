package book

import (
	"database/sql"
	"log"
	"rest-api-postgres/models"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) Service {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FetchBooks() (*[]models.Book, error) {
	var results []models.Book

	sqlStatement := `SELECT * FROM "Book"`

	rows, err := r.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		book := models.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}

func (r *Repository) FetchBook(id int) (*models.Book, error) {
	var book models.Book

	sqlStatement := `SELECT * FROM "Book" WHERE "ID" = $1`

	err := r.db.QueryRow(sqlStatement, id).Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *Repository) InserBook(book models.Book) (*models.Book, error) {
	sqlStatement := `
		INSERT INTO "Book" ("Title", "Author", "Desc", "CreatedAt")
			VALUES($1, $2, $3, $4)
		`
	res, err := r.db.Exec(sqlStatement, book.Title, book.Author, book.Desc, time.Now())
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *Repository) UpdateBook(id int, book models.Book) (*models.Book, error) {
	sqlStatement := `
		UPDATE "Book"
			SET "Title" = $2, "Author" = $3, "Desc" = $4, "UpdatedAt" = $5
		WHERE "ID" = $1
		`
	res, err := r.db.Exec(sqlStatement, id, book.Title, book.Author, book.Desc, time.Now())
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *Repository) RemoveBook(id int) error {
	sqlStatement := `
		UPDATE "Book"
			SET "DeletedAt" = $2
		WHERE "ID" = $1
		`
	res, err := r.db.Exec(sqlStatement, id, time.Now())
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

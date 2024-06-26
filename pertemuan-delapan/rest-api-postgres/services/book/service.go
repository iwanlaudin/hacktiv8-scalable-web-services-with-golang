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

	sql := `SELECT "id", "title", "author_id", "description", "created_at", "updated_at" 
			FROM "book"`

	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		book := models.Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.AuthorId, &book.Description, &book.CreatedAt, &book.UpdatedAt)
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

	sql := `SELECT "id", "title", "author_id", "description", "created_at", "updated_at" 
			FROM "book" 
			WHERE "id" = $1`

	err := r.db.QueryRow(sql, id).Scan(&book.ID, &book.Title, &book.AuthorId, &book.Description)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *Repository) InserBook(book models.Book) (*models.Book, error) {
	sql := `
		INSERT INTO "book" ("title", "author_id", "description", "published", "published_at", "created_at")
			VALUES($1, $2, $3, $4, $5, $6)`

	res, err := r.db.Exec(sql, book.Title, book.AuthorId, book.Description, time.Now())
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
	sql := `
		UPDATE "Book"
		SET "Title" = $2, "Author" = $3, "Desc" = $4, "UpdatedAt" = $5
		WHERE "ID" = $1`

	res, err := r.db.Exec(sql, id, book.Title, book.AuthorId, book.Description, time.Now())
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
	sql := `UPDATE "Book" SET "DeletedAt" = $2 WHERE "ID" = $1`

	res, err := r.db.Exec(sql, id, time.Now())
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

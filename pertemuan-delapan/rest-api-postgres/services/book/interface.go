package book

import "rest-api-postgres/models"

type Service interface {
	FetchBooks() (*[]models.Book, error)
	FetchBook(Id int) (*models.Book, error)
	InserBook(book models.Book) (*models.Book, error)
	UpdateBook(id int, book models.Book) (*models.Book, error)
	RemoveBook(id int) error
}

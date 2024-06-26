package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	AuthorId    string    `json:"author_di"`
	Description string    `json:"description"`
	Published   string    `json:"published"`
	PublishedAt string    `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

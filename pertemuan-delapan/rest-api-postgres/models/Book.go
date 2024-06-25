package models

import "time"

type Book struct {
	ID        int    `json:"book_id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Desc      string `json:"desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

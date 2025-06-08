package models

import (
	"context"
	"perpustakita/internal/db"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func CreateBook(book Book) error {
	_, err := db.DB.Exec(context.Background(),
		"INSERT INTO books (title, author, quantity) VALUES ($1, $2, $3)",
		book.Title, book.Author, book.Quantity)
	return err
}

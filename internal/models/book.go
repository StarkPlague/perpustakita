package models

import (
	"context"
	"log"
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

	if err != nil {
		log.Println("Database error", err)
	}
	return err
}

func GetAllBooks() ([]Book, error) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, title, author, quantity FROM books")
	if err != nil {
		log.Println("Failed to fetch book", err)
		return nil, err
	}
	defer rows.Close()

	var books []Book

	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity)
		if err != nil {
			log.Println("Error scanning row", err)
			return nil, err
		}
		books = append(books, b)
		continue // Skip buku yang gagal di-scan

	}
	return books, nil
}

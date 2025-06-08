package services

import (
	"errors"
	"perpustakita/internal/models"
)

func AddBook(title string, author string, quantity int) error {
	if title == "" || author == "" || quantity <= 0 {
		return errors.New("invalid input")
	}
	book := models.Book{Title: title, Author: author, Quantity: quantity}
	return models.CreateBook(book)

}

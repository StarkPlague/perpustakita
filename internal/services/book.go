package services

import (
	"errors"
	"perpustakita/internal/models"
)

func AddBook(title string, author string, quantity int) error {
	if title == "" {
		return errors.New("null title")
	}
	if author == "" {
		return errors.New("null author")
	}
	if quantity <= 0 {
		return errors.New("quantity cannot be lower than 0")
	}
	
	book := models.Book{Title: title, Author: author, Quantity: quantity}
	return models.CreateBook(book)
}

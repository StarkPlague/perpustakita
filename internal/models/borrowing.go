package models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Borrowing struct {
	ID              int
	BorrowerName    string
	BorrowerContact string
	BookID          string
	BorrowerDate    string
	Status          string
}

func InsertBorrowing(ctx context.Context, conn *pgx.Conn, b Borrowing) error {
	// insert borrowing record
	_, err := conn.Exec(ctx, `
	INSERT INTO borrowings (borrower_name, borrower_contact, book_id)
	values ($1,$2,$3)`,
		b.BorrowerName, b.BorrowerContact, b.BookID)

	if err != nil {
		return err
	}

	//reduce stock
	_, err = conn.Exec(ctx, `
	UPDATE books SET quantity = quantity - 1
	where id = $1 AND quantity > 0`,
		b.BookID)

	return err

}

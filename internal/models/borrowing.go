package models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Borrowing struct {
	ID              int
	BorrowerName    string
	BorrowerContact string
	BookID          int
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

func FindOrCreateBorrower(ctx context.Context, conn *pgx.Conn, name string, contact string) (int, error) {
	var id int
	err := conn.QueryRow(ctx, "SELECT id from borrowers WHERE name = $1 AND contact = $2", name, contact).Scan(&id)
	if err == nil {
		return id, nil 
	}

	err = conn.QueryRow(ctx,
	"INSERT INTO borrowers(name, contact) VALUES ($1, $2) RETURNING id",
	name, contact).Scan(&id)
	return id, err
}

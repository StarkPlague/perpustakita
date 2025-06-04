package models

import (
	"context"

	"perpustakita/internal/db"
)

type Borrower struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	NIK   string `json:"nik"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func GetAllBorrowers() ([]Borrower, error) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, name, nik, email, phone FROM borrowers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrowers []Borrower
	for rows.Next() {
		var b Borrower
		err := rows.Scan(&b.ID, &b.Name, &b.NIK, &b.Email, &b.Phone)
		if err != nil {
			return nil, err
		}
		borrowers = append(borrowers, b)
	}

	return borrowers, nil
}

func InsertBorrower(b Borrower) error {
	_, err := db.DB.Exec(context.Background(), "INSERT INTO borrowers(name, nik, email, phone) VALUES($1, $2, $3, $4)",
		b.Name, b.NIK, b.Email, b.Phone)
	return err
}

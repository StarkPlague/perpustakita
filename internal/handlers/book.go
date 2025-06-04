package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"fmt"
	"perpustakita/internal/db"
	"perpustakita/internal/models"
)

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	author := r.FormValue("author")
	quantityStr := r.FormValue("quantity")

	var quantity int
	_, err := fmt.Sscan(quantityStr, &quantity)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(context.Background(),
		"INSERT INTO books (title, author, quantity) VALUES ($1, $2, $3)",
		title, author, quantity,
	)
	if err != nil {
		fmt.Println("Insert error: ", err)
		http.Error(w, "Failed inserting books", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(context.Background(), "SELECT id, title, author, quantity FROM books")
	if err != nil {
		http.Error(w, "Failed to fetch book", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var b models.Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity)
		if err != nil {
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
		}
		books = append(books, b)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

	id := r.FormValue("id")
	_, err := db.DB.Exec(context.Background(), "DELETE FROM books WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func editBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	title := r.FormValue("title")
	author := r.FormValue("author")
	quantityStr := r.FormValue("quantity")

	var quantity int
	_, err := fmt.Sscan(quantityStr, &quantity)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(context.Background(),
		"UPDATE books SET title=$1, author=$2, quantity=$3 WHERE id=$4",
		title, author, quantity, id,
	)

	if err != nil {
		http.Error(w, "Failed to Update book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func insertDummyBook() {
	_, err := db.DB.Exec(context.Background(),
		"INSERT INTO books (title, author, quantity) VALUES  ($1, $2, $3)",
		"Atomic Habits", "James Clear", 5,
	)
	if err != nil {
		fmt.Println("insert error: ", err)
	} else {
		fmt.Println("New dummy books inserted")
	}
}
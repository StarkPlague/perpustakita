package handlers

import (
	"context"
	"net/http"
	"perpustakita/internal/db"
	"perpustakita/internal/models"
	"strconv"
)

func BorrowBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	contact := r.FormValue("contact")
	bookID := r.FormValue("book_id")

	borrowing := models.Borrowing{
		BorrowerName:    name,
		BorrowerContact: contact,
	}

	//konversi string ke int
	var err error
	borrowing.BookID, err = strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}

	conn, err := db.DB.Acquire(context.Background())
	if err != nil {
		http.Error(w, "failed to acquire database connection", http.StatusInternalServerError)
		return
	}
	defer conn.Release()

	err = models.InsertBorrowing(context.Background(), conn.Conn(), borrowing)
	if err != nil {
		http.Error(w, "failed to borrow book", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

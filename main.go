package main

import (
	"fmt"
	"html/template"
	"net/http"

	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"perpustakita/internal/db"
	"perpustakita/internal/handlers"
)

func init() { //fungsi init otomatis dibaca go sebagai fungsi yang mulai duluan sebelum main
	dsn := "postgres://postgres:1212@localhost:5432/perpustakita"
	db.InitDB(dsn)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)                                                         //index.html
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) //handle static
	http.HandleFunc("/add-book", addBookHandler)
	http.HandleFunc("/delete-book", deleteBookHandler)
	http.HandleFunc("/books", getBookHandler)
	http.HandleFunc("/update-book", editBookHandler)
	//	insertDummyBook()
	fmt.Println("Server Running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

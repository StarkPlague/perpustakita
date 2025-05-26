package main

import (
	"fmt"
	"html/template"
	"net/http"

	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func connectDB() {
	var err error
	db, err = pgx.Connect(context.Background(), "postgres://postgres:1212@localhost:5432/perpustakita")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to DB!")
}

func insertDummyBook() {
	_, err := db.Exec(context.Background(),
		"INSERT INTO books (title, author, quantity) VALUES  ($1, $2, $3)",
		"Atomic Habits", "James Clear", 5,
	)
	if err != nil {
		fmt.Println("insert error: ", err)
	} else {
		fmt.Println("New dummy books inserted")
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")

		if err != nil {
			fmt.Println(err)
			http.Error(w, "error parsing template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})
	connectDB()
	insertDummyBook()
	fmt.Println("Server Running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

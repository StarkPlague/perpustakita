package main

import (
	"log"
	"net/http"

	"github.com/StarkPlague/perpustakita/db"
	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("p!"))
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}

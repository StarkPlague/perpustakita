package handlers

import (
	"encoding/json"
	"net/http"
	"perpustakita/models"
)

func getBorrowers(w http.ResponseWriter, r *http.Request) {
	borrowers, err := models.getAllBorrowers()
	if err != nil {
		http.Error(w, "Failed to get borrowers", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(borrowers)
}

func addBorrower(w http.ResponseWriter, r *http.Request) {
	var b models.Borrower
	err := json.NewEncoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := models.insertBorrower(b)

	if err != nil {
		http.Error(w, "Failed to insert", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

}

package handlers

import (
	"encoding/json"
	"net/http"
	"perpustakita/internal/models"
)

func GetBorrowers(w http.ResponseWriter, r *http.Request) {
	borrowers, err := models.GetAllBorrowers()
	if err != nil {
		http.Error(w, "Failed to get borrowers", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(borrowers)
}

func AddBorrower(w http.ResponseWriter, r *http.Request) {
	var b models.Borrower
	err := json.NewDecoder(r.Body).Decode(&b)

	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = models.InsertBorrower(b)
	if err != nil {
		http.Error(w, "Failed to insert", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

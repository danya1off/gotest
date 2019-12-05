package services

import (
	"encoding/json"
	"github.com/danya1off/library/models"
	"net/http"
)

func GetAllBooks(db *models.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := db.GetAllBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = json.NewEncoder(w).Encode(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

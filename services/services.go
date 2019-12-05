package services

import (
	"encoding/json"
	"github.com/danya1off/library/models"
	"net/http"
)

var booksDao *models.DB

func InitDB(db *models.DB) {
	booksDao = db
}

func GetAllBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := booksDao.GetAllBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = json.NewEncoder(w).Encode(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

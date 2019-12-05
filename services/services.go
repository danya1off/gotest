package services

import (
	"encoding/json"
	"github.com/danya1off/library/utils"
	"net/http"
)

func GetAllBooks(env *utils.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := env.DB.GetAllBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = json.NewEncoder(w).Encode(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

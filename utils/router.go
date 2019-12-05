package utils

import (
	"github.com/danya1off/library/models"
	"github.com/danya1off/library/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Env struct {
	db *models.DB
}

// SetupHTTP : setup routes and run the server on port :3000
func SetupHTTP(db *models.DB) {
	env := Env{db:db}
	r := mux.NewRouter()
	r.HandleFunc("/books", services.GetAllBooks(env.db)).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}

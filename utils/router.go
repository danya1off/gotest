package utils

import (
	"github.com/danya1off/library/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupHTTP : setup routes and run the server on port :3000
func SetupHTTP() {
	r := mux.NewRouter()
	r.HandleFunc("/books", services.GetAllBooks()).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}

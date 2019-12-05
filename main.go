package main

import (
	"github.com/danya1off/library/models"
	"github.com/danya1off/library/services"
	"github.com/danya1off/library/utils"
	"log"
)
func main() {
	db, err := models.CreateDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	services.InitDB(db)
	utils.SetupHTTP()
}

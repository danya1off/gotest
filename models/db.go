package models

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

const (
	postgresURL = "POSTGRES_URL"
	postgres    = "postgres"
)

// DB wrapper
type DB struct {
	*sql.DB
}

// CreateDBConnection creates new connection with database
func CreateDBConnection() (*DB, error) {
	gotenv.Load()
	pqURL, err := pq.ParseURL(os.Getenv(postgresURL))
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(postgres, pqURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("DB created")
	return &DB{db}, nil
}

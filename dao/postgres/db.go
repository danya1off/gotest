package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	DSN string
}

// Connect creates new connection with database
func Connect(cfg *Config) (*sql.DB, error) {
	dbConn, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}
	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

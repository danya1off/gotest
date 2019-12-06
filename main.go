package main

import (
	"net/http"
	"os"

	"github.com/danya1off/library/dao/postgres"
	"github.com/danya1off/library/dao/postgres/book"
	"github.com/danya1off/library/handler"
	"github.com/danya1off/library/service/library"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

const (
	EnvPostgresUrl = "POSTGRES_URL"
)

func main() {
	if err := gotenv.Load(); err != nil {
		logrus.WithError(err).Fatal("Failed to load .env")
	}
	dbCfg := &postgres.Config{
		DSN: os.Getenv(EnvPostgresUrl),
	}
	db, err := postgres.Connect(dbCfg)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to DB")
	}

	bookDao := book.New(db)
	libSvc := library.New(bookDao)
	handlers := handler.New(libSvc)

	http.HandleFunc("/books", handlers.GetAllBooks)
	logrus.Info("Starting http server on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		logrus.WithError(err).Fatal("Failed to start http server")
	}
}

package dao

import "context"

// BookDAO is the DAO interface
type BookDAO interface {
	GetAllBooks(ctx context.Context) ([]Book, error)
}

// Book model for the book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

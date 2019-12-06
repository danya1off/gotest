package book

import (
	"context"
	"database/sql"

	"github.com/danya1off/library/dao"
)

type Dao struct {
	db *sql.DB
}

func New(db *sql.DB) *Dao {
	return &Dao{
		db: db,
	}
}

func (d *Dao) GetAllBooks(ctx context.Context) ([]dao.Book, error) {
	var books []dao.Book
	rows, err := d.db.QueryContext(ctx, "select * from books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book dao.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, err
}

package models

// BookDAO is the DAO interface
type BookDAO interface {
	GetAllBooks() ([]Book, error)
}

// Book : model for the book
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func (db *DB) GetAllBooks() ([]Book, error) {
	var book Book
	var books []Book
	rows, err := db.Query("select * from books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, err
}

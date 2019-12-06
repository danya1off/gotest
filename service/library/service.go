package library

import (
	"context"

	"github.com/danya1off/library/dao"
)

type Service struct {
	bookDao dao.BookDAO
}

func New(bookDao dao.BookDAO) *Service {
	return &Service{bookDao: bookDao}
}

func (s *Service) GetAllBooks(ctx context.Context) ([]dao.Book, error) {
	// any business logic
	return s.bookDao.GetAllBooks(ctx)
}

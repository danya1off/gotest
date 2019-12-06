package service

import (
	"context"

	"github.com/danya1off/library/dao"
)

type Library interface {
	GetAllBooks(ctx context.Context) ([]dao.Book, error)
}

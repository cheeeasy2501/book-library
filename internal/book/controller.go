package book

import (
	"context"
)

type BookControllerInterface interface {
	GetBooks(ctx context.Context, page uint64, limit uint64) ([]Book, error)
}

type BookController struct {
	BookRepo *BookRepository
}

func NewBookController() *BookController {
	return &BookController{
		BookRepo: &BookRepository{},
	}
}

func (bc *BookController) GetBooks(ctx context.Context, page uint64, limit uint64) ([]Book, error) {
	books, err := bc.BookRepo.Get(ctx, page, limit)
	if err != nil {
		return nil, err
	}
	return books, nil
}

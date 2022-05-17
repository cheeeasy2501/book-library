package book

import (
	"context"
)

type BookControllerInterface interface {
	GetBooks(ctx context.Context, page uint64, limit uint64) ([]Book, error)
	GetBook(ctx context.Context, id uint64) (*Book, error)
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

func (bc *BookController) GetBook(ctx context.Context, id uint64) (*Book, error) {
	book, err := bc.BookRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

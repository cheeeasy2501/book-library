package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
)

type BookService struct {
	repo repository.BookRepoInterface
}

func NewBookService(repo repository.BookRepoInterface) *BookService {
	return &BookService{
		repo: repo,
	}
}
func (bs *BookService) Create(ctx context.Context, book *model.Book) error {
	book, err := bs.repo.Create(ctx, book)
	if err != nil {
		return err
	}
	return nil
}
func (bs *BookService) GetAll(ctx context.Context, query model.GetBooksParams) ([]model.Book, error) {
	var books []model.Book
	return books, nil
}
func (bs *BookService) GetById(ctx context.Context, bookId uint64) (*model.Book, error) {
	book, err := bs.repo.GetById(ctx, bookId)
	if err != nil {
		return nil, err
	}

	return book, nil
}
func (bs *BookService) Delete(ctx context.Context, bookId uint64) error {
	return nil
}
func (bs *BookService) Update(ctx context.Context, bookId uint64, input model.Book) error {
	return nil
}

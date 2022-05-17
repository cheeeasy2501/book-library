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
func (bs *BookService) Create(ctx context.Context, book model.Book) (int, error) {
	return 0, nil
}
func (bs *BookService) GetAll(ctx context.Context, query model.GetBooksParams) ([]model.Book, error) {
	var books []model.Book
	return books, nil
}
func (bs *BookService) GetById(ctx context.Context, bookId int) (model.Book, error) {
	var book model.Book
	return book, nil
}
func (bs *BookService) Delete(ctx context.Context, bookId int) error {
	return nil
}
func (bs *BookService) Update(ctx context.Context, bookId int, input model.Book) error {
	return nil
}

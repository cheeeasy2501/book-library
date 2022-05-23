package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
	"time"
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
	currentTime := time.Now()
	book.CreatedAt = currentTime
	book.UpdatedAt = currentTime
	err := bs.repo.Create(ctx, book)
	if err != nil {
		return err
	}
	return nil
}
func (bs *BookService) GetAll(ctx context.Context, paginator forms.Pagination, relationships model.Relationships) ([]model.Book, error) {
	var books []model.Book
	books, err := bs.repo.GetPage(ctx, paginator)
	if err != nil {
		return nil, err
	}

	return books, nil
}
func (bs *BookService) GetById(ctx context.Context, bookId uint64) (*model.Book, error) {
	book, err := bs.repo.GetById(ctx, bookId)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (bs *BookService) Update(ctx context.Context, book *model.Book) error {
	book.UpdatedAt = time.Now()
	err := bs.repo.Update(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookService) Delete(ctx context.Context, bookId uint64) error {
	err := bs.repo.Delete(ctx, bookId)
	if err != nil {
		return err
	}
	return nil
}

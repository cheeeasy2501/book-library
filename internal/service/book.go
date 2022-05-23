package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
	"time"
)

type BookService struct {
	bookRepo   repository.BookRepoInterface
	authorRepo repository.AuthorRepoInterface
}

func NewBookService(bookRepo repository.BookRepoInterface, authorRepo repository.AuthorRepoInterface) *BookService {
	return &BookService{
		bookRepo:   bookRepo,
		authorRepo: authorRepo,
	}
}

func (bs *BookService) Create(ctx context.Context, book *model.Book) error {
	currentTime := time.Now()
	book.CreatedAt = currentTime
	book.UpdatedAt = currentTime
	err := bs.bookRepo.Create(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookService) GetAll(ctx context.Context, paginator forms.Pagination, relations forms.Relations) ([]model.Book, error) {
	var (
		err   error
		books []model.Book
	)

	books, err = bs.bookRepo.GetPage(ctx, paginator, relations)
	if err != nil {
		ctx.Done()
	}

	if len(relations) > 1 {
		books, err = bs.WithRelations(ctx, books, relations)
		if err != nil {
			return nil, err
		}
	}

	return books, nil
}
func (bs *BookService) GetById(ctx context.Context, bookId uint64) (*model.Book, error) {
	book, err := bs.bookRepo.GetById(ctx, bookId)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (bs *BookService) Update(ctx context.Context, book *model.Book) error {
	book.UpdatedAt = time.Now()
	err := bs.bookRepo.Update(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookService) Delete(ctx context.Context, bookId uint64) error {
	err := bs.bookRepo.Delete(ctx, bookId)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookService) WithRelations(ctx context.Context, books []model.Book, relations forms.Relations) ([]model.Book, error) {
	for _, relation := range relations {
		switch relation {
		case forms.Author:
			booksWithAuthors, err := bs.authorRepo.GetBookRelations(ctx, books)
			if err != nil {
				return nil, err
			}

			return booksWithAuthors, nil
		}
	}

	return books, nil
}

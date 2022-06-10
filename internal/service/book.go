package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/relationships"
	"github.com/cheeeasy2501/book-library/internal/repository"
	"time"
)

type BookService struct {
	db               *database.Database
	bookRepository   repository.BookRepoInterface
	authorRepository repository.AuthorRepoInterface
}

func NewBookService(db *database.Database, book repository.BookRepoInterface, author repository.AuthorRepoInterface) *BookService {
	return &BookService{
		db:               db,
		bookRepository:   book,
		authorRepository: author,
	}
}

func (s *BookService) GetAll(ctx context.Context, paginator forms.Pagination, relations relationships.Relations) ([]model.Book, error) {
	var (
		err     error
		bookIds []uint64
		books   []model.Book
		authors model.Authors
	)

	books, err = s.bookRepository.GetPage(ctx, paginator, relations)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(books); i++ {

	}

	authors, err = s.authorRepository.GetAuthorsByBooksIds(ctx, bookIds)
	if err != nil {
		return nil, err
	}

	//for  {
	//
	//}

	return books, nil
}

func (s *BookService) GetById(ctx context.Context, bookId uint64, relations relationships.Relations) (*model.Book, error) {
	book, err := s.bookRepository.GetById(ctx, bookId, relations)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) Create(ctx context.Context, book *model.Book) error {
	currentTime := time.Now()
	book.CreatedAt = currentTime
	book.UpdatedAt = currentTime
	err := s.bookRepository.Create(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) Update(ctx context.Context, book *model.Book) error {
	book.UpdatedAt = time.Now()
	err := s.bookRepository.Update(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookService) Delete(ctx context.Context, bookId uint64) error {
	err := s.bookRepository.Delete(ctx, bookId)
	if err != nil {
		return err
	}
	return nil
}

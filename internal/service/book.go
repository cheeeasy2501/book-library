package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
	"time"
)

type BookServiceInterface interface {
	GetAll(ctx context.Context, params forms.Pagination) ([]model.Book, error)
	GetById(ctx context.Context, bookId uint64) (model.Book, error)
	Create(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, bookId uint64) error

	GetAllWithRelations(ctx context.Context, paginator forms.Pagination) ([]model.FullBook, error)
	GetByIdWithRelations(ctx context.Context, bookId uint64) (model.FullBook, error)
	CreateWithRelations(ctx context.Context, createBook model.CreateBook) (*model.FullBook, error)
}

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

func (s *BookService) GetAll(ctx context.Context, paginator forms.Pagination) ([]model.Book, error) {
	var (
		err   error
		books []model.Book
	)

	books, err = s.bookRepository.GetPage(ctx, paginator.Offset, paginator.Limit)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (s *BookService) GetById(ctx context.Context, bookId uint64) (model.Book, error) {
	book, err := s.bookRepository.GetById(ctx, bookId)
	if err != nil {
		return book, err
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

func (s *BookService) GetAllWithRelations(ctx context.Context, paginator forms.Pagination) ([]model.FullBook, error) {
	var (
		err       error
		books     []model.Book
		fullBooks []model.FullBook
		authors   model.Authors
	)

	books, err = s.bookRepository.GetPage(ctx, paginator.Offset, paginator.Limit)
	if err != nil {
		return nil, err
	}

	mapBooks := make(map[uint64]*model.FullBook, len(books))
	bookIds := make([]uint64, len(books))
	for _, book := range books {
		bookIds = append(bookIds, book.Id)
		mapBooks[book.Id] = &model.FullBook{
			Book: book,
		}
	}

	authors, err = s.authorRepository.GetAuthorsByBooksIds(ctx, bookIds)
	if err != nil {
		return nil, err
	}

	for _, author := range authors {
		book := mapBooks[author.Id]
		book.Authors = append(book.Authors, author)
	}

	return fullBooks, nil
}

func (s *BookService) GetByIdWithRelations(ctx context.Context, bookId uint64) (model.FullBook, error) {
	var (
		err      error
		fullBook model.FullBook
	)

	book, err := s.bookRepository.GetById(ctx, bookId)
	if err != nil {
		return fullBook, err
	}
	fullBook.Book = book
	// implement relations
	return fullBook, nil
}

func (s *BookService) CreateWithRelations(ctx context.Context, m model.CreateBook) (*model.FullBook, error) {
	var (
		err      error
		fullBook model.FullBook
	)

	ctx, finish, err := s.db.TxSession(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		finish(err)
	}()

	err = s.bookRepository.Create(ctx, &m.Book)
	if err != nil {
		return nil, err
	}

	err = s.authorRepository.AttachByAuthorIds(ctx, m.Id, m.AuthorIds)
	if err != nil {
		return nil, err
	}
	authors, err := s.authorRepository.GetAuthorsByBookId(ctx, m.Id)
	if err != nil {
		return &fullBook, err
	}

	currentTime := time.Now()
	fullBook.Book = m.Book
	fullBook.Book.CreatedAt = currentTime
	fullBook.Book.UpdatedAt = currentTime
	fullBook.Authors = authors

	return &fullBook, nil
}

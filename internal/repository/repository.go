package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
)

var builder = sq.StatementBuilderType{}.PlaceholderFormat(sq.Dollar)

type UserRepoInterface interface {
	GetPage(ctx context.Context, paginator forms.Pagination) ([]model.User, error)
	GetById(ctx context.Context, id uint64) (*model.User, error)
	Create(ctx context.Context, usr *model.User) error
	Update(ctx context.Context, usr *model.User) error
	Delete(ctx context.Context, id uint64) error
	FindByUserName(cxt context.Context, username string) (*model.User, error)
}

type AuthorBooksRepoInterface interface {
	Attach(ctx context.Context, authorId, bookId uint64) error
	Detach(ctx context.Context, authorId, bookId uint64) error
}

// TODO add all interfaces for repo there
type Repository struct {
	User        UserRepoInterface
	Book        BookRepoInterface
	Author      AuthorRepoInterface
	AuthorBooks AuthorBooksRepoInterface
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{
		User:        NewUserRepository(db),
		Book:        NewBookRepository(db),
		Author:      NewAuthorRepository(db),
		AuthorBooks: NewAuthorBooksRepository(db),
	}
}

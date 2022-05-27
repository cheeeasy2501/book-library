package repository

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
)

type UserRepoInterface interface {
	GetPage(ctx context.Context, paginator forms.Pagination) ([]model.User, error)
	GetById(ctx context.Context, id uint64) (*model.User, error)
	Create(ctx context.Context, usr *model.User) error
	Update(ctx context.Context, usr *model.User) error
	Delete(ctx context.Context, id uint64) error
	FindByUserName(cxt context.Context, username string) (*model.User, error)
	CheckSignIn(context.Context, *forms.Credentials) (*model.User, error)
}

type BookRepoInterface interface {
	GetPage(ctx context.Context, paginator forms.Pagination) ([]model.Book, error)
	GetById(ctx context.Context, id uint64) (*model.Book, error)
	Create(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id uint64) error
}

//TODO: CHECK IT
type BookAggregateRepoInterface interface {
	GetPage(ctx context.Context, paginator forms.Pagination, relations model.Relations) ([]model.BookAggregate, error)
	GetById(ctx context.Context, id uint64, relations model.Relations) (*model.BookAggregate, error)
}

type AuthorRepoInterface interface {
	GetPage(ctx context.Context, paginator forms.Pagination, relations model.Relations) ([]model.Author, error)
	GetById(ctx context.Context, id uint64) (*model.Author, error)
	Create(ctx context.Context, book *model.Author) error
	Update(ctx context.Context, book *model.Author) error
	Delete(ctx context.Context, id uint64) error
}

// TODO add all interfaces for repo there
type Repository struct {
	User          UserRepoInterface
	Book          BookRepoInterface
	BookAggregate BookAggregateRepoInterface
	Author        AuthorRepoInterface
}

func NewRepository(db *nap.DB) *Repository {
	return &Repository{
		User:          NewUserRepository(db),
		Book:          NewBookRepository(db),
		BookAggregate: NewBookAggregateRepository(db),
		Author:        NewAuthorRepository(db),
	}
}

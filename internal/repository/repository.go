package repository

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
)

type UserRepoInterface interface {
	GetPage(ctx context.Context, page uint64, limit uint64) ([]model.User, error)
	GetById(ctx context.Context, id uint64) (*model.User, error)
	Create(ctx context.Context, usr *model.User) error
	Update(ctx context.Context, usr *model.User) error
	Delete(ctx context.Context, id uint64) error
	FindByUsername(cxt context.Context, username string) (*model.User, error)
	CheckSignIn(context.Context, *model.User) (*model.User, error)
}

type BookRepoInterface interface {
	GetPage(ctx context.Context, page uint64, limit uint64) ([]model.Book, error)
	GetById(ctx context.Context, id uint64) (*model.Book, error)
	Create(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id uint64) error
}

// TODO add all interfaces for repo there
type Repository struct {
	Book BookRepoInterface
	User UserRepoInterface
}

func NewRepository(db *nap.DB) *Repository {
	return &Repository{
		Book: NewBookRepository(db),
		User: NewUserRepository(db),
	}
}

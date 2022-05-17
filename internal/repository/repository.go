package repository

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/google/uuid"
	"github.com/tsenart/nap"
)

type UserRepoInterface interface {
	CheckSignIn(context.Context, *model.User) (*model.User, error)
	Get(context.Context, *model.User)
	FindByUsername(context.Context, string) (*model.User, error)
}

type BookRepoInterface interface {
	GetById(ctx context.Context, id uint64) (*model.Book, error)
	Get(ctx context.Context, page uint64, limit uint64) ([]model.Book, error)
	Create(*model.Book) (*model.Book, error)
	Update(book *model.Book) (*model.Book, error)
	Delete(id uuid.UUID) error
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

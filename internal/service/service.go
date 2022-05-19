package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
)

type AuthorizationServiceInterface interface {
	GenerateToken(usr *model.User) (string, error)
	ParseToken(accessToken string) (int64, error)
	HashPassword(password string) (string, error)
	SignIn(ctx context.Context, usr *model.User) (*model.User, string, error)
	SignUp(ctx context.Context, usr *model.User) (string, error)
}

type UserServiceInterface interface {
	Create(ctx context.Context, user model.User) (uint64, error)
	GetAll(ctx context.Context) ([]model.User, error)
	GetById(ctx context.Context, userId uint64) (model.User, error)
	Delete(ctx context.Context, userId uint64) error
	Update(ctx context.Context, userId uint64) error
}

type BookServiceInterface interface {
	Create(ctx context.Context, book *model.Book) error
	GetAll(ctx context.Context, query model.GetBooksParams) ([]model.Book, error)
	GetById(ctx context.Context, bookId uint64) (*model.Book, error)
	Delete(ctx context.Context, bookId uint64) error
	Update(ctx context.Context, bookId uint64, input model.Book) error
}

type Service struct {
	Authorization AuthorizationServiceInterface
	Book          BookServiceInterface
	User          UserServiceInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.User, "Secret"),
		Book:          NewBookService(repos.Book),
		User:          NewUserService(repos.User),
	}
}

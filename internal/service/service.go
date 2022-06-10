package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/relationships"
	"github.com/cheeeasy2501/book-library/internal/repository"
)

type AuthorizationServiceInterface interface {
	GenerateToken(usr *model.User) (string, error)
	ParseToken(accessToken string) (int64, error)
	HashPassword(password string) ([]byte, error)
	SignIn(ctx context.Context, credentials *forms.Credentials) (*model.User, string, error)
	SignUp(ctx context.Context, userForm *forms.CreateUser) (*model.User, string, error)
}

type UserServiceInterface interface {
	GetAll(ctx context.Context, params forms.Pagination) ([]model.User, error)
	GetById(ctx context.Context, userId uint64) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, userId uint64) error
}

type BookServiceInterface interface {
	GetAll(ctx context.Context, params forms.Pagination, relations relationships.Relations) ([]model.Book, error)
	GetById(ctx context.Context, bookId uint64, relations relationships.Relations) (*model.Book, error)
	Create(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, bookId uint64) error
}

type AuthorServiceInterface interface {
	GetAll(ctx context.Context, params forms.Pagination, relations relationships.Relations) ([]model.Author, error)
	GetById(ctx context.Context, bookId uint64, relations relationships.Relations) (*model.Author, error)
	Create(ctx context.Context, book *model.Author) error
	Update(ctx context.Context, book *model.Author) error
	Delete(ctx context.Context, bookId uint64) error
}

type Service struct {
	Authorization AuthorizationServiceInterface
	User          UserServiceInterface
	Book          BookServiceInterface
	Author        AuthorServiceInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repos.User, "Secret"),
		User:          NewUserService(repos.User),
		Book:          NewBookService(repos.Book),
		Author:        NewAuthorService(repos.Author),
	}
}

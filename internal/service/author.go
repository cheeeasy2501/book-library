package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
)

type AuthorService struct {
	repo repository.AuthorRepoInterface
}

func NewAuthorService(repo repository.AuthorRepoInterface) *AuthorService {
	return &AuthorService{
		repo: repo,
	}
}

func (s *AuthorService) GetAll(ctx context.Context, params forms.Pagination) ([]model.Author, error) {
	return nil, nil
}
func (s *AuthorService) GetById(ctx context.Context, authorId uint64) (*model.Author, error) {
	return nil, nil
}
func (s *AuthorService) Create(ctx context.Context, author *model.Author) error {
	return nil
}
func (s *AuthorService) Update(ctx context.Context, author *model.Author) error {
	return nil
}
func (s *AuthorService) Delete(ctx context.Context, authorId uint64) error {
	return nil
}

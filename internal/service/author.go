package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/relationships"
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

func (s *AuthorService) GetAll(ctx context.Context, pagination forms.Pagination, relations relationships.Relations) ([]model.Author, error) {
	authors, err := s.repo.GetPage(ctx, pagination, relations)
	if err != nil {
		return nil, err
	}
	return authors, nil
}
func (s *AuthorService) GetById(ctx context.Context, authorId uint64, relations relationships.Relations) (*model.Author, error) {
	author, err := s.repo.GetById(ctx, authorId, relations)
	if err != nil {
		return nil, err
	}
	return author, nil
}
func (s *AuthorService) Create(ctx context.Context, author *model.Author) error {
	err := s.repo.Create(ctx, author)
	if err != nil {
		return err
	}

	return nil
}
func (s *AuthorService) Update(ctx context.Context, author *model.Author) error {
	err := s.Update(ctx, author)
	if err != nil {
		return err
	}

	return nil
}
func (s *AuthorService) Delete(ctx context.Context, authorId uint64) error {
	err := s.Delete(ctx, authorId)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
)

type BookAggregateService struct {
	repo repository.BookAggregateRepoInterface
}

func NewBookAggregateService(repo repository.BookAggregateRepoInterface) *BookAggregateService {
	return &BookAggregateService{
		repo: repo,
	}
}

func (s *BookAggregateService) GetAll(ctx context.Context, params forms.Pagination, relationships model.Relations) ([]model.BookAggregate, error) {
	books, err := s.repo.GetPage(ctx, params, relationships)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookAggregateService) GetById(ctx context.Context, bookId uint64, relations model.Relations) (*model.BookAggregate, error) {
	book, err := s.repo.GetById(ctx, bookId, relations)
	if err != nil {
		return nil, err
	}
	return book, nil
}

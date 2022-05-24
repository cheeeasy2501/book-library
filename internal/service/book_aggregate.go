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

func (s *BookAggregateService) GetAll(ctx context.Context, params forms.Pagination, relationships forms.Relations) ([]model.BookAggregate, error) {
	page, err := s.repo.GetPage(ctx, params, relationships)
	if err != nil {
		return nil, err
	}
	return page, nil
}

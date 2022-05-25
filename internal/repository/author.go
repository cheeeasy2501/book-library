package repository

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
)

type Author struct {
	db *nap.DB
}

//TODO: REALIZATION
func NewAuthorRepository(db *nap.DB) *Author {
	return &Author{
		db: db,
	}
}

func (a *Author) GetPage(ctx context.Context, paginator forms.Pagination, relations forms.Relations) ([]model.Author, error) {
	return nil, nil
}
func (a *Author) GetById(ctx context.Context, id uint64) (*model.Author, error) {
	return nil, nil
}
func (a *Author) Create(ctx context.Context, author *model.Author) error {
	return nil
}
func (a *Author) Update(ctx context.Context, author *model.Author) error {
	return nil
}
func (a *Author) Delete(ctx context.Context, id uint64) error {
	return nil
}

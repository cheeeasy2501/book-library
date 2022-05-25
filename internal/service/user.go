package service

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
)

type UserService struct {
	repo repository.UserRepoInterface
}

func NewUserService(repo repository.UserRepoInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) GetAll(ctx context.Context, paginator forms.Pagination) ([]model.User, error) {
	users, err := us.repo.GetPage(ctx, paginator)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) GetById(ctx context.Context, userId uint64) (*model.User, error) {
	user, err := us.repo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) Create(ctx context.Context, user *model.User) error {
	err := us.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) Update(ctx context.Context, user *model.User) error {
	err := us.repo.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) Delete(ctx context.Context, userId uint64) error {
	err := us.repo.Delete(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

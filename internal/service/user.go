package service

import (
	"context"
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

func (us *UserService) Create(ctx context.Context, user model.User) (uint64, error) {
	return 0, nil
}
func (us *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	return users, nil
}
func (us *UserService) GetById(ctx context.Context, userId uint64) (model.User, error) {
	var user model.User
	return user, nil
}
func (us *UserService) Delete(ctx context.Context, userId uint64) error {
	return nil
}
func (us *UserService) Update(ctx context.Context, userId uint64) error {
	return nil
}

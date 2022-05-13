package user

import (
	"cheeeasy2501/book-library/internal/database"
	"context"
)

type UserRepoInterface interface {
	Get(context.Context, *User)
}
type UserRepo struct {
}

func (ur *UserRepo) Get(ctx context.Context, user *User) {
	user, err := database.Instance.Conn.QueryContext()
}

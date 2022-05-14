package user

import (
	"context"
)

type UserRepoInterface interface {
	GetAuth(ctx context.Context, username string, password string) (*User, error)
	Get(context.Context, *User)
}
type UserRepo struct {
}

// GetAuth TODO: mock. replace to repo realization
func (ur *UserRepo) GetAuth(ctx context.Context, username string, password string) (*User, error) {
	usr := &User{
		UserName:  username,
		Password:  password,
		FirstName: "Mock Firstname",
		LastName:  "Mock Lastname",
	}
	return usr, nil
}
func (ur *UserRepo) Get(ctx context.Context, user *User) {
	//user, err := database.Instance.Conn.QueryContext()
}

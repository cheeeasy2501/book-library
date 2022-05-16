package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	e "github.com/cheeeasy2501/book-library/internal/app/errors"
	"github.com/cheeeasy2501/book-library/internal/database"
	"strings"
	"time"
)

type UserRepoInterface interface {
	CheckSignIn(context.Context, *User) (*User, error)
	Get(context.Context, *User)
	FindByUsernameFindByUsername(context.Context, string) (*User, error)
}
type UserRepo struct {
}

// GetAuth TODO: mock. replace to repo realization
func (ur *UserRepo) CheckSignIn(ctx context.Context, usr *User) (*User, error) {
	// return errors.Unauthorized

	return usr, nil
}

func (ur *UserRepo) Get(ctx context.Context, user *User) {
	//user, err := database.Instance.Conn.QueryContext()
}

func (ur *UserRepo) FindByUsername(ctx context.Context, username string) (*User, error) {
	if strings.TrimSpace(username) == "" {
		return nil, e.ValidateError("Username is empty!")
	}

	users := sq.Select("id, firstname, lastname, email, username").From("users")
	query, args, err := users.Where(sq.Eq{"username": username}).PlaceholderFormat(sq.Dollar).ToSql()
	stmt, err := database.Instance.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(args...)
	usr := &User{}
	err = row.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.UserName)

	if err != nil {
		return nil, err
	}

	return usr, err
}

func (ur *UserRepo) Create(ctx context.Context, usr *User) error {
	currentDateTime := time.Now().Format(time.RFC3339)
	query, args, err := sq.Insert("users").Columns("firstname", "lastname", "email", "username", "password", "created_at", "updated_at").
		Values(usr.FirstName, usr.LastName, usr.Email, usr.UserName, usr.Password, currentDateTime, currentDateTime).
		Suffix(`RETURNING "id"`).PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return err
	}

	row, err := database.Instance.Conn.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return err
	}

	usr.Id = id

	return nil
}

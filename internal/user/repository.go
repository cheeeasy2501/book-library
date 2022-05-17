package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserRepoInterface interface {
	CheckSignIn(context.Context, *User) (*User, error)
	Get(context.Context, *User)
	FindByUsername(context.Context, string) (*User, error)
}
type UserRepo struct {
}

// CheckSignIn Check user and password into database
func (ur *UserRepo) CheckSignIn(ctx context.Context, usr *User) (*User, error) {
	find, err := ur.FindByUsername(ctx, usr.UserName)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(find.Password), []byte(usr.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, errors.Unauthorized("Unauthorized user!")
		}

		return nil, err
	}

	return find, err
}

func (ur *UserRepo) Get(ctx context.Context, user *User) {
	//user, err := database.Instance.Conn.QueryContext()
}

// FindByUsername
func (ur *UserRepo) FindByUsername(ctx context.Context, username string) (*User, error) {
	if strings.TrimSpace(username) == "" {
		return nil, errors.ValidateError("Username is empty!")
	}

	users := sq.Select("id, firstname, lastname, email, username, password").From("users")
	query, args, err := users.Where(sq.Eq{"username": username}).PlaceholderFormat(sq.Dollar).ToSql()
	stmt, err := database.Instance.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(args...)
	usr := &User{}
	err = row.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.UserName, &usr.Password)

	if err != nil {
		return nil, err
	}

	return usr, err
}

func (ur *UserRepo) Create(ctx context.Context, usr *User) error {
	var id int64
	currentDateTime := time.Now().Format(time.RFC3339)
	query, args, err := sq.Insert("users").Columns("firstname", "lastname", "email", "username", "password", "created_at", "updated_at").
		Values(usr.FirstName, usr.LastName, usr.Email, usr.UserName, usr.Password, currentDateTime, currentDateTime).
		Suffix(`RETURNING "id"`).PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return err
	}

	err = database.Instance.Conn.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return err
	}
	usr.Id = id

	return nil
}

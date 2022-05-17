package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/errors"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type UserRepository struct {
	db *nap.DB
}

func NewUserRepository(db *nap.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CheckSignIn Check user and password into database
func (ur *UserRepository) CheckSignIn(ctx context.Context, usr *model.User) (*model.User, error) {
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

func (ur *UserRepository) Get(ctx context.Context, user *model.User) {
	//user, err := database.Instance.Conn.QueryContext()
}

// FindByUsername
func (ur *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
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
	usr := &model.User{}
	err = row.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.UserName, &usr.Password)

	if err != nil {
		return nil, err
	}

	return usr, err
}

func (ur *UserRepository) Create(ctx context.Context, usr *model.User) error {
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

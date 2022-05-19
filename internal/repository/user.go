package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

const (
	usersTableName = "users"
)

type UserRepository struct {
	db *nap.DB
}

func NewUserRepository(db *nap.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur UserRepository) GetPage(ctx context.Context, page uint64, limit uint64) ([]model.User, error) {
	var (
		err   error
		users []model.User
	)

	offset := limit * (page - 1)
	query, args, err := sq.Select("id, firstname, lastname, email, username, password, created_at, updated_at").
		From(usersTableName).Limit(limit).Offset(offset).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		usr := model.User{}
		err = rows.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, usr)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return users, nil

}
func (ur *UserRepository) GetById(ctx context.Context, id uint64) (*model.User, error) {
	var (
		err error
		usr *model.User
	)
	query, args, err := sq.Select("id, firstname, lastname, email, username, password, created_at, updated_at").
		From(usersTableName).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, apperrors.UserNotFound
	}

	return usr, nil
}

func (ur *UserRepository) Create(ctx context.Context, usr *model.User) error {
	var id int64
	currentDateTime := time.Now().Format(time.RFC3339)
	query, args, err := sq.Insert(usersTableName).Columns("firstname", "lastname", "email", "username", "password", "created_at", "updated_at").
		Values(usr.FirstName, usr.LastName, usr.Email, usr.UserName, usr.Password, currentDateTime, currentDateTime).
		Suffix(`RETURNING "id"`).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).Scan(&id)
	if err != nil {
		return err
	}
	usr.Id = id

	return nil
}

func (ur *UserRepository) Update(ctx context.Context, usr *model.User) error {
	query, args, err := sq.Update(usersTableName).Set("firstname", usr.FirstName).
		Set("lastname", usr.LastName).Set("email", usr.Email).Set("username", usr.UserName).
		Set("updated_at", usr.UpdatedAt).PlaceholderFormat(sq.Dollar).Suffix("RETURNING created_at").
		Where(sq.Eq{"id": usr.Id}).ToSql()
	if err != nil {
		return err
	}

	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(args...)
	err = result.Scan(&usr.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
func (ur *UserRepository) Delete(ctx context.Context, id uint64) error {
	query, args, err := sq.Delete(usersTableName).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return apperrors.UserNotFound
	}

	return nil
}

// FindByUsername
func (ur *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var (
		err error
		usr *model.User
	)
	if strings.TrimSpace(username) == "" {
		return nil, apperrors.ValidateError("Username is empty!")
	}

	users := sq.Select("id, firstname, lastname, email, username, password, created_at, updated_at").From(usersTableName)
	query, args, err := users.Where(sq.Eq{"username": username}).PlaceholderFormat(sq.Dollar).ToSql()
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(args...)
	err = row.Scan(&usr.Id, &usr.FirstName, &usr.LastName, &usr.Email, &usr.UserName, &usr.Password, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return usr, err
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
			return nil, apperrors.Unauthorized("Unauthorized user!")
		}

		return nil, err
	}

	return find, err
}

package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
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

func (ur UserRepository) GetPage(ctx context.Context, paginator forms.Pagination) ([]model.User, error) {
	var (
		err      error
		users    []model.User
		password string
	)

	query, args, err := builder.
		Select(`id, firstname, lastname, email, username, password, created_at, updated_at`).
		From(usersTableName).
		Offset(paginator.GetOffset()).
		Limit(paginator.Limit).
		ToSql()
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
		user := model.User{}
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		user.SetPassword(password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (ur *UserRepository) GetById(ctx context.Context, id uint64) (*model.User, error) {
	var (
		err      error
		user     *model.User
		password string
	)
	query, args, err := builder.
		Select("id, firstname, lastname, email, username, password, created_at, updated_at").
		From(usersTableName).Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).
		Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, apperrors.UserNotFound
	}
	user.SetPassword(password)

	return user, nil
}

func (ur *UserRepository) Create(ctx context.Context, user *model.User) error {
	var id int64
	currentDateTime := time.Now()
	query, args, err := builder.
		Insert(usersTableName).
		Columns("firstname", "lastname", "email", "username", "password", "created_at", "updated_at").
		Values(
			user.FirstName,
			user.LastName,
			user.Email,
			user.UserName,
			user.Password(),
			currentDateTime,
			currentDateTime,
		).
		Suffix(`RETURNING "id"`).
		ToSql()
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
	user.Id = id
	user.CreatedAt = currentDateTime
	user.UpdatedAt = currentDateTime

	return nil
}

func (ur *UserRepository) Update(ctx context.Context, usr *model.User) error {
	query, args, err := builder.
		Update(usersTableName).
		Set("firstname", usr.FirstName).
		Set("lastname", usr.LastName).
		Set("email", usr.Email).
		Set("username", usr.UserName).
		Set("updated_at", usr.UpdatedAt).
		Suffix("RETURNING created_at").
		Where(sq.Eq{"id": usr.Id}).
		ToSql()
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
	query, args, err := builder.
		Delete(usersTableName).
		Where(sq.Eq{"id": id}).
		ToSql()
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

// FindByUserName FindByUsername
func (ur *UserRepository) FindByUserName(ctx context.Context, username string) (*model.User, error) {
	var (
		err      error
		password string
	)

	user := &model.User{}
	query, args, err := builder.Select("id, firstname, lastname, email, username, password, created_at, updated_at").
		From(usersTableName).
		Where(sq.Eq{"username": username}).
		ToSql()
	stmt, err := ur.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).
		Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.UserName,
			&password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	user.SetPassword(password)

	return user, err
}

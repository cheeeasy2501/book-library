package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
)

const (
	bookTableName = "books"
)

type BookRepository struct {
	db *nap.DB
}

func NewBookRepository(db *nap.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) GetPage(ctx context.Context, paginator forms.Pagination) ([]model.Book, error) {
	var (
		err   error
		books []model.Book
	)
	query, args, err := builder.
		Select(`books.id, books.house_publish_id, books.title, books.description, 
		books.link, books.in_stock, books.created_at, books.updated_at`).
		From(bookTableName).
		Limit(paginator.Limit).
		Offset(paginator.GetOffset()).
		ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := br.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := model.Book{}
		err = rows.Scan(
			&book.Id,
			&book.HousePublishId,
			&book.Title,
			&book.Description,
			&book.Link,
			&book.InStock,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (br *BookRepository) GetById(ctx context.Context, id uint64) (*model.Book, error) {
	var book model.Book
	query, args, err := builder.
		Select(`books.id, books.house_publish_id, books.title, books.description, 
		books.link, books.in_stock, books.created_at, books.updated_at`).
		From(bookTableName).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := br.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).
		Scan(
			&book.Id,
			&book.HousePublishId,
			&book.Title,
			&book.Description,
			&book.Link,
			&book.InStock,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, apperrors.BookNotFound
	}

	return &book, nil
}

func (br *BookRepository) Create(ctx context.Context, book *model.Book) error {
	query, args, err := builder.
		Insert(bookTableName).
		Columns(`house_publish_id, title, description, link, in_stock, created_at, updated_at`).
		Values(
			book.Title,
			book.Description,
			book.Link,
			book.InStock,
			book.CreatedAt,
			book.UpdatedAt,
		).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()
	if err != nil {
		return err
	}

	stmt, err := br.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(args...)
	err = result.Scan(
		&book.Id,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (br *BookRepository) Update(ctx context.Context, book *model.Book) error {
	query, args, err := builder.
		Update(bookTableName).
		Set("house_publish_id", book.HousePublishId).
		Set("title", book.Title).
		Set("description", book.Description).
		Set("link", book.Link).
		Set("updated_at", book.UpdatedAt).
		Suffix("RETURNING created_at").
		Where(sq.Eq{"id": book.Id}).
		ToSql()
	if err != nil {
		return err
	}

	stmt, err := br.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(args...)
	err = result.Scan(&book.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (br *BookRepository) Delete(ctx context.Context, id uint64) error {
	query, args, err := builder.
		Delete(bookTableName).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	stmt, err := br.db.PrepareContext(ctx, query)
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
		return apperrors.BookNotFound
	}

	return nil
}

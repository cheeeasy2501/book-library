package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/model"
	"time"
)

const (
	BookTableName = "books"
)

type BookRepoInterface interface {
	GetPage(ctx context.Context, offset, limit uint64) ([]model.Book, error)
	GetById(ctx context.Context, id uint64) (model.Book, error)
	Create(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id uint64) error
}

type BookRepository struct {
	db *database.Database
}

func NewBookRepository(db *database.Database) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetPage(ctx context.Context, offset, limit uint64) ([]model.Book, error) {
	var (
		books []model.Book
	)

	b := builder.Select(`books.id, books.house_publish_id, books.title, 
	     books.description, books.link, books.in_stock, books.created_at, books.updated_at`).
		From(BookTableName)

	query, args, err := b.
		Offset(offset).
		Limit(limit).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

	return books, nil
}

func (r *BookRepository) GetById(ctx context.Context, id uint64) (model.Book, error) {
	var (
		err error
	)

	book := model.Book{}

	b := builder.Select(`books.id, books.house_publish_id, books.title, 
	     books.description, books.link, books.in_stock, books.created_at, books.updated_at`).
		From(BookTableName).
		Where(sq.Eq{"books.id": id})

	query, args, err := b.GroupBy("books.id").
		ToSql()
	if err != nil {
		return book, err
	}

	err = r.db.QueryRowContext(ctx, query, args...).
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
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *BookRepository) Create(ctx context.Context, book *model.Book) error {
	tx := r.db.GetTxSession(ctx)
	now := time.Now()
	query, args, err := builder.
		Insert(BookTableName).
		Columns("house_publish_id", "title", "description", "link", "in_stock, created_at, updated_at").
		Values(
			book.HousePublishId,
			book.Title,
			book.Description,
			book.Link,
			book.InStock,
			now,
			now,
		).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()
	if err != nil {
		return err
	}

	err = tx.QueryRowContext(ctx, query, args...).
		Scan(
			&book.Id,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Update(ctx context.Context, book *model.Book) error {
	query, args, err := builder.
		Update(BookTableName).
		Set("house_publish_id", book.HousePublishId).
		Set("title", book.Title).
		Set("description", book.Description).
		Set("link", book.Link).
		Set("updated_at", book.UpdatedAt).
		Suffix("RETURNING updated_at").
		Where(sq.Eq{"id": book.Id}).
		ToSql()
	if err != nil {
		return err
	}

	result := r.db.QueryRowContext(ctx, query, args...)
	err = result.Scan(&book.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Delete(ctx context.Context, id uint64) error {
	query, args, err := builder.
		Delete(BookTableName).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, query, args...)
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

package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
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

func (br *BookRepository) GetPage(ctx context.Context, page uint64, limit uint64) ([]model.Book, error) {
	var (
		err   error
		books []model.Book
	)
	offset := limit * (page - 1)
	query, args, err := sq.Select("id, author_id, title, description, link, in_stock, created_at, updated_at").From(bookTableName).Limit(limit).Offset(offset).
		PlaceholderFormat(sq.Dollar).ToSql()
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
		err = rows.Scan(&book.Id, &book.AuthorId, &book.Title, &book.Description, &book.Link, &book.InStock, &book.CreatedAt, &book.UpdatedAt)
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
	query, args, err := sq.Select("id, author_id, title, description, link, in_stock, created_at, updated_at").From(bookTableName).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := br.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).Scan(&book.Id, &book.AuthorId, &book.Title, &book.Description, &book.Link, &book.InStock, &book.CreatedAt, &book.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, apperrors.BookNotFound
	}

	return &book, nil
}

func (br *BookRepository) Create(ctx context.Context, book *model.Book) error {
	err := book.Validate()
	if err != nil {
		return err
	}

	query, args, err := sq.Insert(bookTableName).Columns("author_id, title, description, link, in_stock, created_at, updated_at").
		Values(book.AuthorId, book.Title, book.Description, book.Link, book.InStock, book.CreatedAt, book.UpdatedAt).PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, created_at, updated_at").ToSql()
	if err != nil {
		return err
	}
	stmt, err := br.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result := stmt.QueryRow(args...)
	err = result.Scan(&book.Id, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (br *BookRepository) Update(ctx context.Context, book *model.Book) error {
	query, args, err := sq.Update(bookTableName).Set("author_id", book.AuthorId).
		Set("title", book.Title).Set("description", book.Description).Set("link", book.Link).
		Set("updated_at", book.UpdatedAt).PlaceholderFormat(sq.Dollar).Suffix("RETURNING created_at").
		Where(sq.Eq{"id": book.Id}).ToSql()
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
	query, args, err := sq.Delete(bookTableName).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
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

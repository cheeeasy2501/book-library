package book

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/database"
	e "github.com/cheeeasy2501/book-library/internal/errors"
	"github.com/google/uuid"
)

type BookInterface interface {
	GetById(id int64) (*Book, error)
	Get(page int, limit int) ([]*Book, error)
	Create(*Book) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(id uuid.UUID) error
}

type BookRepository struct {
}

func (br *BookRepository) Get(ctx context.Context, page uint64, limit uint64) ([]Book, error) {
	var (
		books []Book
	)
	offset := limit * (page - 1)
	query, args, err := sq.Select("*").From("books").Limit(limit).Offset(offset).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	stmt, err := database.Instance.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := Book{}
		err = rows.Scan(&book.ID, &book.AuthorID, &book.Title, &book.Description, &book.Link, &book.InStock, &book.CreatedAt, &book.UpdatedAt)
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

func (br *BookRepository) GetById(ctx context.Context, id uint64) (*Book, error) {
	var book Book

	query, args, err := sq.Select("*").From("books").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := database.Instance.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).Scan(&book.ID, &book.AuthorID, &book.Title, &book.Description, &book.Link, &book.InStock, &book.CreatedAt, &book.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, e.NotFoundError("Book isn't found")
	}

	return &book, nil
}

func (br *BookRepository) Create(book *Book) (*Book, error) {

	return book, nil
}

func (br *BookRepository) Update(book *Book) (*Book, error) {

	return book, nil
}

func (br *BookRepository) Delete(id uuid.UUID) error {

	return nil
}

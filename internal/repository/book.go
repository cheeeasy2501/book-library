package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/database"
	e "github.com/cheeeasy2501/book-library/internal/errors"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/google/uuid"
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

func (br *BookRepository) Get(ctx context.Context, page uint64, limit uint64) ([]model.Book, error) {
	var (
		books []model.Book
	)
	offset := limit * (page - 1)
	query, args, err := sq.Select("*").From(bookTableName).Limit(limit).Offset(offset).
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
		book := model.Book{}
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

func (br *BookRepository) GetById(ctx context.Context, id uint64) (*model.Book, error) {
	var book model.Book

	query, args, err := sq.Select("*").From(bookTableName).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
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

func (br *BookRepository) Create(book *model.Book) (*model.Book, error) {

	return book, nil
}

func (br *BookRepository) Update(book *model.Book) (*model.Book, error) {

	return book, nil
}

func (br *BookRepository) Delete(id uuid.UUID) error {

	return nil
}

package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	e "github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/google/uuid"
	"github.com/tsenart/nap"
	"time"
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

func (br *BookRepository) GetByPage(ctx context.Context, page uint64, limit uint64) ([]model.Book, error) {
	var (
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

	query, args, err := sq.Select("id, author_id, title, description, link, in_stock, created_at, updated_at").From(bookTableName).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := br.db.PrepareContext(ctx, query)
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

func (br *BookRepository) Create(ctx context.Context, book *model.Book) (*model.Book, error) {
	err := book.Validate()
	if err != nil {
		return nil, err
	}
	currentTime := time.Now()
	query, args, err := sq.Insert(bookTableName).Columns("author_id, title, description, link, in_stock, created_at, updated_at").
		Values(book.AuthorID, book.Title, book.Description, book.Link, book.InStock, currentTime, currentTime).PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, created_at, updated_at").ToSql()
	if err != nil {
		return nil, err
	}
	stmt, err := br.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	result := stmt.QueryRow(args...)
	err = result.Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (br *BookRepository) Update(book *model.Book) (*model.Book, error) {

	return book, nil
}

func (br *BookRepository) Delete(id uuid.UUID) error {

	return nil
}

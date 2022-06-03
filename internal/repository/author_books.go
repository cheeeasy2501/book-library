package repository

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/tsenart/nap"
)

// Work with auhtor_books table
// Attach and detach authors or books
// Delete where author_id = 1 or book_id = 1

const (
	authorBooksTableName = "author_books"
)

type AuthorBooks struct {
	db *nap.DB
}

func NewAuthorBooksRepository(db *nap.DB) *AuthorBooks {
	return &AuthorBooks{
		db: db,
	}
}

func (ab *AuthorBooks) Attach(ctx context.Context, authorId, bookId uint64) error {
	query, args, err := builder.Insert(authorBooksTableName).
		Columns("author_id", "book_id").
		Values(authorId, bookId).
		ToSql()
	prepareContext, err := ab.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	result, err := prepareContext.Exec(args...)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("row affected is zero")
	}

	return nil
}

func (ab *AuthorBooks) Detach(ctx context.Context, authorId, bookId uint64) error {
	query, args, err := builder.Delete(authorBooksTableName).
		Where(sq.Eq{
			"author_id": authorId,
			"book_id":   bookId,
		}).
		ToSql()

	prepareContext, err := ab.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	result, err := prepareContext.Exec(args...)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("row affected is zero")
	}

	return nil
}

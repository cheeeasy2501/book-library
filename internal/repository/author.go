package repository

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/relationships"
	"golang.org/x/exp/slices"
)

const (
	AuthorTableName = "authors"
)

type AuthorRepoInterface interface {
	GetPage(ctx context.Context, paginator forms.Pagination, relations relationships.Relations) ([]model.FullAuthor, error)
	GetById(ctx context.Context, id uint64, relations relationships.Relations) (*model.FullAuthor, error)
	Create(ctx context.Context, book *model.FullAuthor) error
	Update(ctx context.Context, book *model.FullAuthor) error
	Delete(ctx context.Context, id uint64) error

	GetAuthorsByBookId(ctx context.Context, id uint64) (model.Authors, error)
	GetAuthorsByBooksIds(ctx context.Context, ids []uint64) (model.Authors, error)
	AttachByAuthorIds(ctx context.Context, bookId uint64, authorIds []uint64) error
}

type Author struct {
	db *database.Database
}

func NewAuthorRepository(db *database.Database) *Author {
	return &Author{
		db: db,
	}
}

func (r *Author) GetAuthorsByBookId(ctx context.Context, bookId uint64) (model.Authors, error) {
	var authors model.Authors
	tx := r.db.GetTxSession(ctx)
	query, args, err := builder.Select(
		"authors.id, authors.firstname, authors.lastname, authors.created_at, authors.updated_at").
		From(authorBooksTableName).
		LeftJoin("authors on authors.id = author_books.author_id").
		Where(sq.Eq{"author_books.book_id": bookId}).
		OrderBy("author_books.book_id").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		author := model.Author{}
		err := rows.Scan(
			&author.Id,
			&author.FirstName,
			&author.LastName,
			&author.CreatedAt,
			&author.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func (r *Author) GetAuthorsByBooksIds(ctx context.Context, bookIds []uint64) (model.Authors, error) {
	var authors model.Authors

	query, args, err := builder.Select(
		`authors.id, authors.firstname, authors.lastname, authors.created_at, authors.updated_at`).
		From(authorBooksTableName).
		LeftJoin("authors on authors.id = author_books.author_id").
		Where(sq.Eq{"author_books.book_id": bookIds}).
		OrderBy("author_books.book_id").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		author := model.Author{}
		err = rows.Scan(
			&author.Id,
			&author.FirstName,
			&author.LastName,
			&author.CreatedAt,
			&author.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}

func (r *Author) GetPage(ctx context.Context, paginator forms.Pagination, relations relationships.Relations) ([]model.FullAuthor, error) {
	var authors []model.FullAuthor
	b := builder.Select("authors.id, authors.firstname, authors.lastname, authors.created_at, authors.updated_at")
	b.From(AuthorTableName).
		Offset(paginator.GetOffset()).
		Limit(paginator.Limit)

	if slices.Contains(relations, relationships.BookRelation) {
		b = b.Columns("json_agg(books.*) as books").
			LeftJoin("author_books on books.id = author_books.book_id").
			LeftJoin("authors on authors.id = author_books.author_id").
			GroupBy("authors.id")
	}
	query, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		author := model.FullAuthor{}
		scan := []interface{}{
			&author.Id,
			&author.FirstName,
			&author.LastName,
			&author.CreatedAt,
			&author.UpdatedAt,
		}
		if slices.Contains(relations, relationships.BookRelation) {
			scan = append(scan, &author.Books)
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (r *Author) GetById(ctx context.Context, id uint64, relations relationships.Relations) (*model.FullAuthor, error) {
	author := &model.FullAuthor{}
	scan := []interface{}{
		&author.Id,
		&author.FirstName,
		&author.LastName,
		&author.CreatedAt,
		&author.UpdatedAt,
	}
	b := builder.Select("authors.id, authors.firstname, authors.lastname, authors.created_at, authors.updated_at")
	b.From(AuthorTableName).
		Where(sq.Eq{"authors.id": id})

	if slices.Contains(relations, relationships.BookRelation) {
		scan = append(scan, []interface{}{
			&author.Books,
		})
		b = b.Columns("json_agg(books.*) as books").
			LeftJoin("author_books on books.id = author_books.book_id").
			LeftJoin("authors on authors.id = author_books.author_id").
			GroupBy("authors.id")
	}
	query, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(scan...)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (r *Author) Create(ctx context.Context, author *model.FullAuthor) error {
	query, args, err := builder.
		Insert(AuthorTableName).
		Columns(`house_publish_id, title, description, link, in_stock, created_at, updated_at`).
		Values(
			author.FirstName,
			author.LastName,
		).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()
	if err != nil {
		return err
	}

	result := r.db.QueryRowContext(ctx, query, args...)
	err = result.Scan(
		&author.Id,
		&author.CreatedAt,
		&author.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Author) Update(ctx context.Context, author *model.FullAuthor) error {
	query, args, err := builder.
		Update(AuthorTableName).
		Set("firstname", author.FirstName).
		Set("lastname", author.LastName).
		Set("created_at", author.CreatedAt).
		Set("updated_at", author.UpdatedAt).
		Suffix("RETURNING updated_at").
		Where(sq.Eq{"id": author.Id}).
		ToSql()
	if err != nil {
		return err
	}
	result := r.db.QueryRowContext(ctx, query, args...)
	err = result.Scan(&author.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *Author) Delete(ctx context.Context, id uint64) error {
	query, args, err := builder.
		Delete(AuthorTableName).
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
		return apperrors.AuthorNotFound
	}

	return nil
}

func (a *Author) AttachByAuthorIds(ctx context.Context, bookId uint64, authorIds []uint64) error {
	var count int64 = 0
	session := a.db.GetTxSession(ctx)
	for _, authorId := range authorIds {
		query, args, err := builder.Insert(authorBooksTableName).
			Columns("author_id", "book_id").
			Values(authorId, bookId).
			ToSql()
		result, err := session.ExecContext(ctx, query, args...)
		if err != nil {
			return err
		}
		affected, err := result.RowsAffected()

		if err != nil {
			return err
		}
		count += affected
	}

	if int(count) != len(authorIds) {
		return errors.New("Invalid updated")
	}

	return nil
}

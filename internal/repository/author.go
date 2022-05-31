package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/relationships"
	"github.com/tsenart/nap"
	"golang.org/x/exp/slices"
)

const (
	AuthorTableName = "author"
)

type Author struct {
	db *nap.DB
}

func NewAuthorRepository(db *nap.DB) *Author {
	return &Author{
		db: db,
	}
}

func (a *Author) GetPage(ctx context.Context, paginator forms.Pagination, relations relationships.Relations) ([]model.Author, error) {
	var authors []model.Author

	b := builder.Select("author.id, author.firstname, author.lastname, author.created_at, author.updated_at")
	b.From(AuthorTableName).
		Offset(paginator.GetOffset()).
		Limit(paginator.Limit)

	if slices.Contains(relations, relationships.BookRelation) {
		b = b.Columns("json_agg(books.*) as books").
			LeftJoin("author_books on books.id = author_books.book_id").
			LeftJoin("author on author.id = author_books.author_id").
			GroupBy("author.id")
	}
	query, args, err := b.ToSql()

	if err != nil {
		return nil, err
	}

	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		author := model.Author{}
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

func (a *Author) GetById(ctx context.Context, id uint64, relations relationships.Relations) (*model.Author, error) {
	author := &model.Author{}
	scan := []interface{}{
		&author.Id,
		&author.FirstName,
		&author.LastName,
		&author.CreatedAt,
		&author.UpdatedAt,
	}
	b := builder.Select("author.id, author.firstname, author.lastname, author.created_at, author.updated_at")
	b.From(AuthorTableName).
		Where(sq.Eq{"author.id": id})

	if slices.Contains(relations, relationships.BookRelation) {
		scan = append(scan, []interface{}{
			&author.Books,
		})
		b = b.Columns("json_agg(books.*) as books").
			LeftJoin("author_books on books.id = author_books.book_id").
			LeftJoin("author on author.id = author_books.author_id").
			GroupBy("author.id")
	}
	query, args, err := b.ToSql()

	if err != nil {
		return nil, err
	}

	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(args...).Scan(scan...)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (a *Author) Create(ctx context.Context, author *model.Author) error {
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

	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(args...)
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

func (a *Author) Update(ctx context.Context, author *model.Author) error {
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

	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(args...)
	err = result.Scan(&author.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (a *Author) Delete(ctx context.Context, id uint64) error {
	query, args, err := builder.
		Delete(AuthorTableName).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	stmt, err := a.db.PrepareContext(ctx, query)
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
		return apperrors.AuthorNotFound
	}

	return nil
}

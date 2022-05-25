package repository

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
)

type BookAggregateRepository struct {
	db *nap.DB
}

func NewBookAggregateRepository(db *nap.DB) *BookAggregateRepository {
	return &BookAggregateRepository{db: db}
}

func (bar *BookAggregateRepository) GetPage(ctx context.Context, paginator forms.Pagination, relations forms.Relations) ([]model.BookAggregate, error) {
	var (
		books []model.BookAggregate
	)
	// Get books with join authors
	query, args, err := sq.Select(fmt.Sprintf("books.*, json_agg(%s)", "author.*")).
		From("books").LeftJoin("author_books on books.id = author_books.book_id").
		LeftJoin("author on author.id = author_books.author_id").
		GroupBy("books.id").OrderBy("books.id").
		Limit(paginator.Limit).Offset(paginator.GetOffset()).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := bar.db.PrepareContext(ctx, query)
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
		book := model.BookAggregate{}
		rows.Scan(&book.Id, &book.PublishHouseId, &book.Title, &book.Description, &book.Link, &book.InStock,
			&book.CreatedAt, &book.UpdatedAt, &book.BookAuthors)
		books = append(books, book)
	}

	return books, nil
}

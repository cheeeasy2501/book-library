package repository

import (
	"context"
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

func (bar *BookAggregateRepository) GetPage(ctx context.Context, paginator forms.Pagination, relations model.Relations) ([]model.BookAggregate, error) {
	var (
		books []model.BookAggregate
	)

	// Get books with join authors and publish house
	query, args, err := sq.Select(`books.id, books.title, books.description, books.link, books.in_stock, books.created_at, books.updated_at, 
		json_agg(author.*) as authors, house_publishes.*`).
		From("books").LeftJoin("author_books on books.id = author_books.book_id").
		LeftJoin("author on author.id = author_books.author_id").
		LeftJoin("house_publishes on books.publishhouse_id = house_publishes.id").
		GroupBy("books.id", "house_publishes.id").OrderBy("books.id").
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
		err = rows.Scan(&book.Id, &book.Title, &book.Description, &book.Link, &book.InStock,
			&book.CreatedAt, &book.UpdatedAt, &book.Relations.BookAuthors, &book.Relations.BookHousePublishes.Id,
			&book.Relations.BookHousePublishes.Name, &book.Relations.BookHousePublishes.CreatedAt, &book.Relations.BookHousePublishes.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (bar *BookAggregateRepository) GetById(ctx context.Context, id uint64, relations model.Relations) (*model.BookAggregate, error) {
	var (
		err error
	)
	//build := builder.NewBookBuilder(ctx, relations)
	//book, err := build.WithAuthors().Execute(bar.db)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return nil, nil
	//return book, nil
}

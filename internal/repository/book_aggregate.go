package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/relationships"
	"github.com/tsenart/nap"
	"golang.org/x/exp/slices"
)

type BookAggregateRepository struct {
	db *nap.DB
}

func NewBookAggregateRepository(db *nap.DB) *BookAggregateRepository {
	return &BookAggregateRepository{db: db}
}

func (bar *BookAggregateRepository) GetPage(ctx context.Context, paginator forms.Pagination, relations relationships.Relations) ([]model.BookAggregate, error) {
	var (
		books []model.BookAggregate
		scan  []interface{}
	)

	b := builder.Select(`books.id, books.house_publish_id, books.title, 
	     books.description, books.link, books.in_stock, books.created_at, books.updated_at 
		`).
		From(bookTableName)
	withAuthors := slices.Contains(relations, relationships.AuthorRel)
	if withAuthors {
		b = b.Columns(`json_agg(author.*) as authors`).
			LeftJoin("author_books on books.id = author_books.book_id").
			LeftJoin("author on author.id = author_books.author_id")
	}
	withPublishHouse := slices.Contains(relations, relationships.PublishHouseRel)
	if withPublishHouse {
		b = b.Columns(`house_publishes.id, house_publishes.name,
			house_publishes.created_at, house_publishes.updated_at`).
			LeftJoin("house_publishes on books.house_publish_id = house_publishes.id").
			GroupBy("house_publishes.id")
	}
	query, args, err := b.
		GroupBy("books.id").
		Offset(paginator.GetOffset()).
		Limit(paginator.Limit).
		ToSql()
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
		scan = []interface{}{
			&book.Id,
			&book.HousePublishId,
			&book.Title,
			&book.Description,
			&book.Link,
			&book.InStock,
			&book.CreatedAt,
			&book.UpdatedAt,
		}

		if withAuthors {
			scan = append(
				scan,
				&book.Relations.BookAuthors,
			)
		}

		if withPublishHouse {
			book.Relations.BookHousePublishes = &model.BookHousePublishes{}
			scan = append(
				scan,
				&book.Relations.BookHousePublishes.Id,
				&book.Relations.BookHousePublishes.Name,
				&book.Relations.BookHousePublishes.CreatedAt,
				&book.Relations.BookHousePublishes.UpdatedAt,
			)
		}

		err = rows.Scan(scan...)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (bar *BookAggregateRepository) GetById(ctx context.Context, id uint64, relations relationships.Relations) (*model.BookAggregate, error) {
	var (
		err error
	)

	book := &model.BookAggregate{}
	scan := []interface{}{
		&book.Id,
		&book.HousePublishId,
		&book.Title,
		&book.Description,
		&book.Link,
		&book.InStock,
		&book.CreatedAt,
		&book.UpdatedAt,
	}

	b := builder.Select(`books.id, books.house_publish_id, books.title, 
	     books.description, books.link, books.in_stock, books.created_at, books.updated_at 
		`).
		From(bookTableName).
		Where(sq.Eq{"books.id": id})

	if slices.Contains(relations, relationships.AuthorRel) {
		book.Relations.BookAuthors = model.BookAuthors{}
		scan = append(
			scan,
			&book.Relations.BookAuthors,
		)

		b = b.Columns(`json_agg(author.*) as authors`).
			LeftJoin("author_books on books.id = author_books.book_id").
			LeftJoin("author on author.id = author_books.author_id")
	}

	if slices.Contains(relations, relationships.PublishHouseRel) {
		book.Relations.BookHousePublishes = &model.BookHousePublishes{}
		scan = append(
			scan,
			&book.Relations.BookHousePublishes.Id,
			&book.Relations.BookHousePublishes.Name,
			&book.Relations.BookHousePublishes.CreatedAt,
			&book.Relations.BookHousePublishes.UpdatedAt,
		)
		b = b.Columns(`house_publishes.id, house_publishes.name,
			house_publishes.created_at, house_publishes.updated_at`).
			LeftJoin("house_publishes on books.house_publish_id = house_publishes.id").
			GroupBy("house_publishes.id")
	}
	query, args, err := b.GroupBy("books.id").
		ToSql()
	if err != nil {
		return nil, err
	}

	stmt, err := bar.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(args...)
	if err != nil {
		return nil, err
	}

	err = row.Scan(scan...)
	if err != nil {
		return nil, err
	}

	return book, nil
}

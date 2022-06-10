package repository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/relationships"
	"golang.org/x/exp/slices"
	"time"
)

const (
	BookTableName = "books"
)

const (
	AuthorBooksTableName = "author_books"
)

type BookRepository struct {
	db *database.Database
}

func NewBookRepository(db *database.Database) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetPage(ctx context.Context, paginator forms.Pagination, relations relationships.Relations) ([]model.Book, error) {
	var (
		books []model.Book
		scan  []interface{}
	)

	b := builder.Select(`books.id, books.house_publish_id, books.title, 
	     books.description, books.link, books.in_stock, books.created_at, books.updated_at`).
		From(BookTableName)
	//withAuthors := slices.Contains(relations, relationships.AuthorRel)
	//if withAuthors {
	//	b = b.Columns(`json_agg(author.*) as authors`).
	//		LeftJoin("author_books on books.id = author_books.book_id").
	//		LeftJoin("author on author.id = author_books.author_id")
	//}
	//withPublishHouse := slices.Contains(relations, relationships.PublishHouseRel)
	//if withPublishHouse {
	//	b = b.Columns(`house_publishes.id, house_publishes.name,
	//		house_publishes.created_at, house_publishes.updated_at`).
	//		LeftJoin("house_publishes on books.house_publish_id = house_publishes.id").
	//		GroupBy("house_publishes.id")
	//}
	query, args, err := b.
		//GroupBy("books.id").
		Offset(paginator.Offset).
		Limit(paginator.Limit).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := model.Book{}
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
				&book.Authors,
			)
		}

		if withPublishHouse {
			book.BookHousePublishes = &model.BookHousePublishes{}
			scan = append(
				scan,
				&book.BookHousePublishes.Id,
				&book.BookHousePublishes.Name,
				&book.BookHousePublishes.CreatedAt,
				&book.BookHousePublishes.UpdatedAt,
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

func (br *BookRepository) GetById(ctx context.Context, id uint64, relations relationships.Relations) (*model.Book, error) {
	var (
		err error
	)

	book := &model.Book{}
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
		From(BookTableName).
		Where(sq.Eq{"books.id": id})

	if slices.Contains(relations, relationships.AuthorRel) {
		book.Authors = model.Authors{}
		scan = append(
			scan,
			&book.Authors,
		)

		b = b.Columns(`json_agg(author.*) as authors`).
			LeftJoin("author_books on books.id = author_books.book_id").
			LeftJoin("author on author.id = author_books.author_id")
	}

	if slices.Contains(relations, relationships.PublishHouseRel) {
		book.BookHousePublishes = &model.BookHousePublishes{}
		scan = append(
			scan,
			&book.BookHousePublishes.Id,
			&book.BookHousePublishes.Name,
			&book.BookHousePublishes.CreatedAt,
			&book.BookHousePublishes.UpdatedAt,
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

	stmt, err := br.db.PrepareContext(ctx, query)
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

func (br *BookRepository) Create(ctx context.Context, book *model.Book) error {
	tx, err := br.GetTx(ctx)
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}(tx)
	now := time.Now()
	query, args, err := builder.
		Insert(BookTableName).
		Columns("house_publish_id", "title", "description", "link", "in_stock, created_at, updated_at").
		Values(
			book.HousePublishId,
			book.Title,
			book.Description,
			book.Link,
			book.InStock,
			now,
			now,
		).
		Suffix("RETURNING id, created_at, updated_at").
		ToSql()
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, args...).Scan(&book.Id, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return err
	}

	if len(book.Authors) != 0 {
		for _, author := range book.Authors {
			query, args, err = builder.Insert(AuthorBooksTableName).
				Columns("author_id", "book_id").
				Values(
					author.Id,
					book.Id,
				).
				ToSql()

			stmt, err = tx.PrepareContext(ctx, query)
			if err != nil {
				return err
			}

			_, err = stmt.Exec(args...)
			if err != nil {
				return err
			}

			query, args, err = builder.
				Select("firstname, lastname, created_at, updated_at").
				From(AuthorTableName).
				Where(sq.Eq{"id": author.Id}).
				ToSql()
			stmt, err = tx.Prepare(query)
			if err != nil {
				return err
			}
			err = stmt.QueryRow(args...).Scan(
				&author.FirstName,
				&author.LastName,
				&author.CreatedAt,
				&author.UpdatedAt,
			)
			if err != nil {
				return err
			}
		}
	}
	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (br *BookRepository) Update(ctx context.Context, book *model.Book) error {
	query, args, err := builder.
		Update(BookTableName).
		Set("house_publish_id", book.HousePublishId).
		Set("title", book.Title).
		Set("description", book.Description).
		Set("link", book.Link).
		Set("updated_at", book.UpdatedAt).
		Suffix("RETURNING updated_at").
		Where(sq.Eq{"id": book.Id}).
		ToSql()
	if err != nil {
		return err
	}

	stmt, err := br.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(args...)
	err = result.Scan(&book.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (br *BookRepository) Delete(ctx context.Context, id uint64) error {
	query, args, err := builder.
		Delete(BookTableName).
		Where(sq.Eq{"id": id}).
		ToSql()
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

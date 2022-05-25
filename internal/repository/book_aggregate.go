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
		//author        model.Author
		books []model.BookAggregate
	)
	//bookMap := make(map[uint64]model.BookAggregate, 0)
	//TODO: REALIZATION FOR BOOK AGGREGATE
	//pq.Array  author.Columns() - not work - author.*
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
		//TODO REALIZATION
		book := model.BookAggregate{}
		var test, mest string
		err = rows.Scan(&book.Id, &mest, &book.Title,
			&book.Description, &book.Link, &book.InStock, &book.CreatedAt, &book.UpdatedAt,
			&test)
		if err != nil {
			return nil, err
		}
	}

	return books, nil
}

//func (a *Author) GetBookRelations(ctx context.Context, books []model.Book) ([]model.Book, error) {
//	var (
//		bookId  uint64
//		bookIds []uint64
//		author  model.Author
//	)
//	for _, book := range books {
//		bookIds = append(bookIds, book.Id)
//	}
//	query, args, err := sq.Select("book_id", author.Columns()).
//		From("book_authors").Join("JOIN author USING (author_id)").
//		Where(sq.Eq{"book_id": bookIds}).PlaceholderFormat(sq.Dollar).ToSql()
//	if err != nil {
//		return nil, err
//	}
//
//	stmt, err := a.db.PrepareContext(ctx, query)
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	rows, err := stmt.Query(args...)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	scan := append(author.Fields(), bookId)
//	for rows.Next() {
//		err = rows.Scan(scan)
//		if err != nil {
//			return nil, err
//		}
//
//		for _, book := range books {
//			if book.Id == bookId {
//				book.Authors = append(book.Authors, author)
//				break
//			}
//		}
//	}
//
//	return books, nil
//}

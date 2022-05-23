package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
)

type Author struct {
	db *nap.DB
}

func NewAuthorRepository(db *nap.DB) *Author {
	return &Author{
		db: db,
	}
}

func (a *Author) GetPage(ctx context.Context, paginator forms.Pagination, relations forms.Relations) ([]model.Author, error) {
	return nil, nil
}
func (a *Author) GetById(ctx context.Context, id uint64) (*model.Author, error) {
	return nil, nil
}
func (a *Author) Create(ctx context.Context, author *model.Author) error {
	return nil
}
func (a *Author) Update(ctx context.Context, author *model.Author) error {
	return nil
}
func (a *Author) Delete(ctx context.Context, id uint64) error {
	return nil
}

func (a *Author) GetBookRelations(ctx context.Context, books []model.Book) ([]model.Book, error) {
	var (
		bookId  uint64
		bookIds []uint64
		author  model.Author
	)
	for _, book := range books {
		bookIds = append(bookIds, book.Id)
	}
	query, args, err := sq.Select("book_id", author.Columns()).
		From("book_authors").Join("JOIN author USING (author_id)").
		Where(sq.Eq{"book_id": bookIds}).PlaceholderFormat(sq.Dollar).ToSql()
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

	scan := append(author.Fields(), bookId)
	for rows.Next() {
		err = rows.Scan(scan)
		if err != nil {
			return nil, err
		}

		for _, book := range books {
			if book.Id == bookId {
				book.Authors = append(book.Authors, author)
				break
			}
		}
	}

	return books, nil
}

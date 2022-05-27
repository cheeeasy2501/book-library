package builder

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/tsenart/nap"
	"golang.org/x/exp/slices"
)

//Обертка для всех Builder
type Builder struct {
	BookBuilder BookBuilderInterface
}

//func NewBuilder(builder sq.SelectBuilder) *Builder {
//	return &Builder{
//		BookBuilder: NewBookBuilder(builder),
//	}
//}

type BasicBuilderInterface interface {
	Columns() []string
	Fields() []interface{}
}

type FieldMap map[string]FieldItem

// ScanFields Поля для scan и SelectFields
type SelectFields []string    // идишники
type JoinStrings []string     // набор join
type ScanFields []interface{} // ссылки в модель

type FieldItem struct {
	SelectFields *SelectFields
	JoinStrings  *JoinStrings
	ScanFields   *ScanFields
}

type BasicBuilder struct {
	Context   context.Context
	Builder   sq.SelectBuilder // составляет sql запросы
	FieldMap  FieldMap         // map с SelectFields - 'book.id,book.title' и с ссылками на поля модели для Scan метода
	Relations model.Relations
}

//Интерфейс BookBuilder
type BookBuilderInterface interface {
	WithAuthors() *BookBuilder
	WithPublishHouse() *BookBuilder
}

type BookBuilder struct {
	BasicBuilder
	model *model.BookAggregate
}

func NewBookBuilder(ctx context.Context, relations model.Relations) *BookBuilder {
	m := &model.BookAggregate{}
	fieldMap := FieldMap{}
	fieldMap["book"] = FieldItem{
		SelectFields: &SelectFields{"books.id", "books.title", "books.description", "books.link", "books.in_stock", "books.created_at", "books.updated_at"},
		ScanFields:   &ScanFields{&m.Id, &m.Title, &m.Description, &m.Link, &m.InStock, &m.CreatedAt, &m.UpdatedAt}}
	builder := sq.SelectBuilder{}.From("books").PlaceholderFormat(sq.Dollar)

	return &BookBuilder{
		BasicBuilder: BasicBuilder{
			Context:   ctx,
			Builder:   builder,
			FieldMap:  fieldMap,
			Relations: relations,
		},
		model: m,
	}
}

func (b *BookBuilder) WithAuthors() *BookBuilder {
	if slices.Contains(b.Relations, model.PublishHouseRel) {
		b.FieldMap["authors"] = FieldItem{
			SelectFields: &SelectFields{"json_agg(author.*) as authors"},
			JoinStrings:  &JoinStrings{"author_books on books.id = author_books.book_id", "author on author.id = author_books.author_id"},
			ScanFields:   &ScanFields{&b.model.Relations.BookAuthors},
		}
		b.Builder = b.Builder.GroupBy("books.id")
	}

	return b
}

//TODO: отвязать от model?
func (b *BookBuilder) Execute(conn *nap.DB) (*model.BookAggregate, error) {
	var scan = &ScanFields{}
	for _, item := range b.FieldMap {
		b.Builder = b.Builder.Columns(*item.SelectFields...)
		if item.JoinStrings != nil && len(*item.JoinStrings) != 0 {
			for _, join := range *item.JoinStrings {
				b.Builder = b.Builder.LeftJoin(join)
			}
		}
		*scan = append(*scan, *item.ScanFields...)
	}

	query, args, err := b.Builder.ToSql()
	stmt, err := conn.PrepareContext(b.Context, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(args...)
	if err != nil {
		return nil, err
	}

	err = row.Scan(*scan...)
	if err != nil {
		return nil, err
	}

	return b.model, nil
}

func (b *BookBuilder) WithPublishHouse() *BookBuilder {
	return b
}

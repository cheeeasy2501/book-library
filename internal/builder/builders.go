package builder

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/tsenart/nap"
)

// Builder Обертка для всех Builder
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

type FieldMap map[string]*FieldItem

type ScanFieldsInterface interface {
	SetScan(string, *FieldItem)
}

type SelectFields []string    // набор Select для запроса
type JoinStrings []string     // набор Join для запроса
type ScanFields []interface{} // слайс с ссылками на поля структуры(куда нужно сетить)

// FieldItem Структура для отношений между queryColumn - scan link
type FieldItem struct {
	SelectFields *SelectFields
	JoinStrings  *JoinStrings
	ScanFields   *ScanFields
}

type BasicBuilder struct {
	Context  context.Context
	Builder  sq.SelectBuilder // составляет sql запросы
	FieldMap FieldMap         // map с SelectFields - 'book.id,book.title' и с ссылками на поля модели для Scan метода
}

type BookBuilderInterface interface {
	WithAuthors() *BookBuilder
	WithPublishHouse() *BookBuilder
}

type BookBuilder struct {
	BasicBuilder
	model ScanFieldsInterface
}

func NewBookBuilder(ctx context.Context, m ScanFieldsInterface) *BookBuilder {
	fieldMap := FieldMap{}
	fieldMap["book"] = &FieldItem{
		SelectFields: &SelectFields{
			"books.id",
			"books.title",
			"books.description",
			"books.link",
			"books.in_stock",
			"books.created_at",
			"books.updated_at",
		},
	}
	m.SetScan("book", fieldMap["book"])
	builder := sq.SelectBuilder{}.From("books").PlaceholderFormat(sq.Dollar)

	return &BookBuilder{
		BasicBuilder: BasicBuilder{
			Context:  ctx,
			Builder:  builder,
			FieldMap: fieldMap,
		},
		model: m,
	}
}

// WithAuthors Relation
func (b *BookBuilder) WithAuthors() *BookBuilder {
	if _, ok := b.FieldMap["authors"]; !ok {
		b.FieldMap["authors"] = &FieldItem{
			SelectFields: &SelectFields{"json_agg(author.*) as authors"},
			JoinStrings: &JoinStrings{
				"author_books on books.id = author_books.book_id",
				"author on author.id = author_books.author_id",
			},
		}
		b.model.SetScan("authors", b.FieldMap["authors"])
		b.Builder = b.Builder.GroupBy("books.id")
	}

	return b
}

// WithPublishHouse Relation
func (b *BookBuilder) WithPublishHouse() *BookBuilder {
	if _, ok := b.FieldMap["publish_house"]; !ok {
		b.FieldMap["publish_house"] = &FieldItem{
			SelectFields: &SelectFields{
				"house_publishes.*",
			},
			JoinStrings: &JoinStrings{
				"house_publishes on books.publishhouse_id = house_publishes.id",
			},
		}
		b.model.SetScan("publish_house", b.FieldMap["publish_house"])
		b.Builder = b.Builder.GroupBy("house_publishes.id")
	}

	return b
}

// Execute Метод перегоняет результат запроса в структуру
func (b *BookBuilder) Execute(conn *nap.DB) error {
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
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(args...)
	if err != nil {
		return err
	}

	err = row.Scan(*scan...)
	if err != nil {
		return err
	}

	return nil
}

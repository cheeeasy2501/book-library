package builder

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
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

type SelectFields []string // набор Select для запроса

type JoinVector string

const (
	LeftJoin  = JoinVector("LEFT")
	RightJoin = JoinVector("RIGHT")
	InnerJoin = JoinVector("INNER")
)

type JoinString struct {
	JoinVector JoinVector
	Table      string
	Condition  string
}

type JoinStrings []JoinString // набор Join для запроса
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

func (b *BasicBuilder) joinToString(j *JoinString) string {
	return fmt.Sprintf("%s %s", j.Table, j.Condition)
}

// Execute Метод перегоняет результат запроса в структуру
func (b *BasicBuilder) Execute(conn *nap.DB) error {
	var scan = &ScanFields{}
	for _, item := range b.FieldMap {
		b.Builder = b.Builder.Columns(*item.SelectFields...)
		if item.JoinStrings != nil && len(*item.JoinStrings) != 0 {
			for _, join := range *item.JoinStrings {
				joinStr := b.joinToString(&join)
				switch join.JoinVector {
				case LeftJoin:
					b.Builder = b.Builder.LeftJoin(joinStr)
				case RightJoin:
					b.Builder = b.Builder.RightJoin(joinStr)
				case InnerJoin:
					b.Builder = b.Builder.InnerJoin(joinStr)
				default:
					logrus.Info("Invalid join closure!")
				}

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

type BookBuilderInterface interface {
	WithAuthors() *BookBuilder
	WithPublishHouse() *BookBuilder
}

type BookBuilder struct {
	BasicBuilder
	Model ScanFieldsInterface
}

func NewBookBuilder(ctx context.Context, m ScanFieldsInterface) *BookBuilder {
	fieldMap := make(FieldMap, 0)
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
		Model: m,
	}
}

// WithAuthors Relation
func (b *BookBuilder) WithAuthors() *BookBuilder {
	if _, ok := b.FieldMap["authors"]; !ok {
		joinStrings := &JoinStrings{
			JoinString{
				JoinVector: LeftJoin,
				Table:      "author_books",
				Condition:  "on books.id = author_books.book_id",
			},
			JoinString{
				JoinVector: LeftJoin,
				Table:      "authors",
				Condition:  "on authors.id = author_books.author_id",
			},
		}

		b.FieldMap["authors"] = &FieldItem{
			SelectFields: &SelectFields{"json_agg(authors.*) as authors"},
			JoinStrings:  joinStrings,
		}

		b.Model.SetScan("authors", b.FieldMap["authors"])
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
				JoinString{
					JoinVector: LeftJoin,
					Table:      "house_publishes",
					Condition:  "on books.house_publish_id = house_publishes.id",
				},
			},
		}
		b.Model.SetScan("publish_house", b.FieldMap["publish_house"])
		b.Builder = b.Builder.GroupBy("house_publishes.id")
	}

	return b
}

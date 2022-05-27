package builder

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/cheeeasy2501/book-library/internal/model"
)

//Обертка для всех Builder
type Builder struct {
	BookBuilder BookBuilderInterface
}

func NewBuilder(builder sq.SelectBuilder) *Builder {
	return &Builder{
		BookBuilder: NewBookBuilder(builder),
	}
}

type BasicBuilderInterface interface {
	Columns() []string
	Fields() []interface{}
	Build() (string, []interface{}, error)
}

// ScanFields Поля для scan и SelectFields
type ScanFields map[string][]interface{} // ссылки в модель
type ColumnFields map[string][]string    // идишники

type BasicBuilder struct {
	builder      sq.SelectBuilder // составляет sql запросы
	ScanFields   *ScanFields
	ColumnFields *ColumnFields
	relations    *model.Relations
}

// билдим  sql-ник и  scan
func (b *BasicBuilder) Build() (string, []interface{}, error) {
	sql, args, err := b.builder.ToSql()
	if err != nil {
		return "", nil, err
	}

	return sql, args, nil
}

//Интерфейс BookBuilder
type BookBuilderInterface interface {
	WithAllRelations() *BookBuilder
	WithAuthors() *BookBuilder
	WithPublishHouse() *BookBuilder
}

type BookBuilder BasicBuilder

// как-то забрать из модели или описать тут
//func (a *BookBuilder) Columns() []string {
//	return []string{"author.id", "author.firstname", "author.lastname", "author.created_at", "author.updated_at"}
//}
//
//func (a *BookBuilder) Fields() []interface{} {
//	return []interface{}{&a.Id, &a.Title, &a.Description, &a.Link, &a.InStock, &a.CreatedAt, &a.UpdatedAt}
//}

func NewBookBuilder(relations *model.Relations) *BookBuilder {
	//sq.SelectBuilder{}
	builder := sq.Select(`books.id, books.title, books.description, books.link, books.in_stock, books.created_at, books.updated_at 
		`).From("books")
	return &BookBuilder{
		builder:   builder,
		relations: relations,
	}
}

func (b *BookBuilder) WithAllRelations() (string, []interface{}, error) {
	return b.WithAuthors().WithPublishHouse().ToSql()
}

func (b *BookBuilder) WithAuthors() *BookBuilder {
	book := model.Book{}
	b.RelationFields
	//b.builder.Columns(book.Columns()...)

	return b
}

func (b *BookBuilder) WithPublishHouse() *BookBuilder {
	return b
}

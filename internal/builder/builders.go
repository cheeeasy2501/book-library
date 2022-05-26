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

//Интерфейс BookBuilder
type BookBuilderInterface interface {
	WithAllRelations() sq.SelectBuilder
	WithAuthors() sq.SelectBuilder
	WithPublishHouse() sq.SelectBuilder
}

//Поля для scan и Relations
type ScanFields map[string][]interface{}
type RelationFields map[string][]string

type BookBuilder struct {
	builder        sq.SelectBuilder
	ScanFields     *ScanFields
	RelationFields *RelationFields
}

// как-то забрать из модели или описать тут
//func (a *BookBuilder) Columns() []string {
//	return []string{"author.id", "author.firstname", "author.lastname", "author.created_at", "author.updated_at"}
//}
//
//func (a *BookBuilder) Fields() []interface{} {
//	return []interface{}{&a.Id, &a.Title, &a.Description, &a.Link, &a.InStock, &a.CreatedAt, &a.UpdatedAt}
//}

func NewBookBuilder(b sq.SelectBuilder) *BookBuilder {
	return &BookBuilder{
		builder: b,
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

func (b *BookBuilder) ToSql() (string, []interface{}, error) {
	sql, args, err := b.builder.ToSql()
	if err != nil {
		return "", nil, err
	}

	return sql, args, nil
}

package model

import (
	"encoding/json"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/lann/builder"
)

type BookAuthors []Author

// impliment sql.Scanner
func (b *BookAuthors) Scan(src interface{}) error {
	bts, ok := src.([]byte)
	if !ok {
		return errors.New("Error Scanning Array")
	}

	return json.Unmarshal(bts, b)
}

type BookAggregate struct {
	Book
	Relations struct { //TODO: check it and refactoring
		BookAuthors        BookAuthors         `json:"authors,omitempty"`
		BookHousePublishes *BookHousePublishes `json:"house_publishes,omitempty"`
	} `json:"relations"`
}

// Relations
const (
	AuthorRel       = Relation("authors")
	PublishHouseRel = Relation("publish_house")
)

func GetBookRelations() []Relation {
	return []Relation{AuthorRel, PublishHouseRel}
}

func (ba *BookAggregate) WithRelations(builder *builder.Builder, scan *[]interface{}, rel *Relationships) {

}

func (ba *BookAggregate) WithAuthors(sb *sq.SelectBuilder, scan *[]interface{}) sq.SelectBuilder {
	*scan = append(*scan, &ba.Relations.BookAuthors) //scanfields
	return sb.Columns(`json_agg(author.*) as authors`).LeftJoin("author_books on books.id = author_books.book_id").
		LeftJoin("author on author.id = author_books.author_id")
}

func (ba *BookAggregate) WithPublishHouse(sb *sq.SelectBuilder, scan *[]interface{}) sq.SelectBuilder {
	ba.Relations.BookHousePublishes = &BookHousePublishes{}
	*scan = append(*scan, &ba.Relations.BookHousePublishes.Id, &ba.Relations.BookHousePublishes.Name,
		&ba.Relations.BookHousePublishes.CreatedAt, &ba.Relations.BookHousePublishes.UpdatedAt)
	return sb.Columns(`house_publishes.*`).LeftJoin("house_publishes on books.publishhouse_id = house_publishes.id").
		GroupBy("house_publishes.id")
}

// TODO: Trying to create Mapper
func (a *BookAggregate) Columns() string {
	//return a.Book.Columns()
	return ""
}

func (a *BookAggregate) Fields() []interface{} {
	return []interface{}{&a.Book.Id, &a.Title, &a.Description, &a.Link, &a.InStock, &a.CreatedAt, a.UpdatedAt}
}

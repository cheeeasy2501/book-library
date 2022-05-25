package model

import (
	"encoding/json"
	"errors"
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
	BookAuthors `json:"authors,omitempty"`
}

func (a *BookAggregate) Columns() string {
	author := Author{}
	columns := a.Book.Columns() + ", " + author.Columns()
	//+ a.Authors[0].Columns()
	return columns
}

func (a *BookAggregate) Fields() []interface{} {
	author := Author{}
	fields := []interface{}{&a.Id, &a.Title, &a.Description, &a.Link, &a.InStock, &a.CreatedAt, a.UpdatedAt}
	return append(fields, author.Fields()...)
}

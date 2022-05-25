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

//type BookHousePublishes string

type BookAggregate struct {
	Book
	BookHousePublishes BookHousePublishes `json:"house_publishes,omitempty"`
	BookAuthors        BookAuthors        `json:"authors,omitempty"`
}

// TODO: Trying to create Mapper
//func (a *BookAggregate) Columns() string {
//	return a.Book.Columns()
//}
//
//func (a *BookAggregate) Fields() []interface{} {
//	return []interface{}{&a.Book.Id, &a.Title, &a.Description, &a.Link, &a.InStock, &a.CreatedAt, a.UpdatedAt}
//}

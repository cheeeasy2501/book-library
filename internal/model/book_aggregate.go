package model

import (
	"encoding/json"
	"errors"
	"github.com/cheeeasy2501/book-library/internal/builder"
)

type BookAuthors []Author

// Scan impliment sql.Scanner
func (b *BookAuthors) Scan(src interface{}) error {
	bts, ok := src.([]byte)
	if !ok {
		return errors.New("Error Scanning Array")
	}

	return json.Unmarshal(bts, b)
}

type BookAggregate struct {
	Book
	Relations struct {
		BookAuthors        BookAuthors        `json:"authors,omitempty"`
		BookHousePublishes BookHousePublishes `json:"house_publishes,omitempty"`
	} `json:"relations"`
}

// SetScan заполняет ScanFields ссылками из структуры
func (ba *BookAggregate) SetScan(key string, fieldItem *builder.FieldItem) {
	switch key {
	case "book":
		fieldItem.ScanFields = &builder.ScanFields{
			&ba.Id,
			&ba.Title,
			&ba.Description,
			&ba.Link,
			&ba.InStock,
			&ba.CreatedAt,
			&ba.UpdatedAt,
		}
	case "authors":
		fieldItem.ScanFields = &builder.ScanFields{
			&ba.Relations.BookAuthors,
		}
	case "publish_house":
		fieldItem.ScanFields = &builder.ScanFields{
			&ba.Relations.BookHousePublishes.Id,
			&ba.Relations.BookHousePublishes.Name,
			&ba.Relations.BookHousePublishes.CreatedAt,
			&ba.Relations.BookHousePublishes.UpdatedAt,
		}
	}
}

// Relations
const (
	AuthorRel       = Relation("authors")
	PublishHouseRel = Relation("publish_house")
)

func GetBookRelations() []Relation {
	return []Relation{AuthorRel, PublishHouseRel}
}

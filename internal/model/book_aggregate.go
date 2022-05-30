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
	Relations struct { //TODO: check it and refactoring
		BookAuthors        BookAuthors         `json:"authors,omitempty"`
		BookHousePublishes *BookHousePublishes `json:"house_publishes,omitempty"`
	} `json:"relations"`
}

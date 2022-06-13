package model

import (
	"encoding/json"
	"errors"
)

type Author struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Timestamp
	Books Books `json:"books,omitempty"`
}

type Books []Book

// impliment sql.Scanner
func (ab *Books) Scan(src interface{}) error {
	bts, ok := src.([]byte)
	if !ok {
		return errors.New("Error Scanning Array")
	}

	return json.Unmarshal(bts, ab)
}

type AuthorAggregate struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Timestamp
	BookId uint64 `json:"bookId"`
}

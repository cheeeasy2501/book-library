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
	Books AuthorBooks `json:"books,omitempty"`
}

type AuthorBooks []Book

// impliment sql.Scanner
func (ab *AuthorBooks) Scan(src interface{}) error {
	bts, ok := src.([]byte)
	if !ok {
		return errors.New("Error Scanning Array")
	}

	return json.Unmarshal(bts, ab)
}

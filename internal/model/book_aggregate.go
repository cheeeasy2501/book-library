package model

import (
	"bytes"
	"encoding/json"
	"errors"
)

type Authors []Author

// impliment sql.Scanner
func (a *Authors) Scan(src interface{}) error {
	bts, ok := src.([]byte)
	if !ok {
		return errors.New("Error Scanning Array")
	}

	return json.Unmarshal(bts, a)
}

func (a *Authors) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	//todo check it
	var authors []Author
	err := json.Unmarshal(data, &authors)
	if err != nil {
		return err
	}
	*a = append(*a, authors...)

	return nil
}

//func (a *Authors) UnmarshalText(data []byte) error {
//	p := strings.Split(string(data), ",")
//
//	for _, value := range p {
//		*a = append(*a, value)
//	}
//
//	return nil
//}

type BookAggregate struct {
	Book
	Relations struct { //TODO: check it and refactoring
		Authors            Authors             `json:"authors,omitempty"`
		BookHousePublishes *BookHousePublishes `json:"house_publishes,omitempty"`
	} `json:"relations"`
}

//type AuthorBooksAggregate struct {
//	Id       uint64
//	AuthorId uint64
//	BookId   uint64
//}

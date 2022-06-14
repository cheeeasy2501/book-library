package model

import (
	"bytes"
	"encoding/json"
	"errors"
)

type Book struct {
	Id             uint64 `json:"id"`
	HousePublishId uint64 `json:"house_publish_id"`
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description"`
	Link           string `json:"link" binding:"url"`
	InStock        uint   `json:"in_stock"`
	Timestamp
}

type FullBook struct {
	Book
	Authors            []AuthorRelation    `json:"authors,omitempty"`
	BookHousePublishes *BookHousePublishes `json:"house_publishes,omitempty"`
}

type Authors []FullAuthor

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
	var authors []FullAuthor
	err := json.Unmarshal(data, &authors)
	if err != nil {
		return err
	}
	*a = append(*a, authors...)

	return nil
}

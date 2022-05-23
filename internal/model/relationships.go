package model

import (
	"bytes"
	"strings"
)

const (
	BookAuthorRelation = Relation("author")
)

type (
	Relation  string
	Relations []Relation
)

func (r *Relations) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	p := strings.Split(string(data), ",")

	for _, value := range p {
		*r = append(*r, Relation(value))
	}

	return nil
}

func (r Relations) MarshalJSON() ([]byte, error) {
	n := len(r)
	if n == 0 {
		return nil, nil
	}

	buff := bytes.NewBuffer(nil)
	for index, value := range r {
		buff.WriteString(string(value))
		if index < n-1 {
			buff.WriteString(",")
		}
	}

	return buff.Bytes(), nil
}

type Relationships struct {
	Relations Relations `json:"relations" form:"relations"`
}

type BooksRelationships struct {
	Relationships
}

func (b BooksRelationships) LoadAndValidate() {
	for index, value := range b.Relations {
		if value != BookAuthorRelation {
			b.Relations = append(b.Relations[:index], b.Relations[index+1:]...)
		}
	}

}

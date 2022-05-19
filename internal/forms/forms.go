package forms

import (
	"bytes"
	"strings"
)

type FormInterface interface {
	LoadAndValidate()
}

type Relationships struct {
	Relations []Relation `json:"relations"`
}

type Relation string

const (
	AuthorRelation Relation = "author"
)

type Relations []Relation

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

type Pagination struct {
	Relationships
	Page  uint64 `form:"page" json:"page" binding:"required,gte=1"`
	Limit uint64 `form:"limit" json:"limit" binding:"required,gte=1"`
}

func (pf *Pagination) LoadAndValidate() {}

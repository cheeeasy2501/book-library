package forms

import (
	"bytes"
	"golang.org/x/exp/slices"
	"strings"
)

//TODO: WORK WITH RELATION.
type (
	Relation      string
	Relations     []Relation
	Relationships struct {
		Relations `form:"relations"`
	}
)

// implements encoding.TextUnmarshaler
func (r *Relations) UnmarshalText(data []byte) error {
	p := strings.Split(string(data), ",")

	for _, value := range p {
		*r = append(*r, Relation(value))
	}

	return nil
}

func (r Relations) MarshalText() ([]byte, error) {
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

func (r *Relations) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	return r.UnmarshalText(data)
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

const (
	AuthorRel Relation = Relation("authors")
	TestRel   Relation = Relation("tests")
)

func (r Relation) String() string {
	return r.String()
}

func GetBookRelations() []Relation {
	return []Relation{AuthorRel, TestRel}
}

func (r *Relations) FilterRelations(relations []Relation) []Relation {
	filteredRelations := []Relation{}
	//TODO: error 0 - authors,test123,tests invalid value
	for _, value := range *r {
		if slices.Contains(relations, value) {
			filteredRelations = append(filteredRelations, value)
		}
	}

	return filteredRelations
}

package forms

import (
	"bytes"
	"golang.org/x/exp/slices"
	"strings"
)

type (
	Relation      string
	Relations     []Relation
	Relationships struct {
		Relations
	}
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

const (
	Author Relation = Relation("author")
	Test   Relation = Relation("test")
)

func (r Relation) String() string {
	return r.String()
}

func (r Relationships) BookRelations() []Relation {
	relations := []Relation{Author, Test}
	for index, value := range r.Relations {
		if !slices.Contains(relations, value) {
			relations = append(relations[:index], relations[index+1:]...)
		}
	}

	return relations
}

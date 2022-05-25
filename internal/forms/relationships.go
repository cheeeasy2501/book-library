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

// implements encoding.TextMarshaler
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

//implement json.Unmarshaler
func (r *Relations) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	return r.UnmarshalText(data)
}

//implement json.Marshaler
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
	AuthorRel       = Relation("authors")
	PublishHouseRel = Relation("publish_house")
)

func (r *Relations) FilterRelations(relations []Relation) []Relation {
	filteredRelations := []Relation{}
	for _, value := range *r {
		if slices.Contains(relations, value) {
			filteredRelations = append(filteredRelations, value)
		}
	}

	return filteredRelations
}

// Relations
func GetBookRelations() []Relation {
	return []Relation{AuthorRel, PublishHouseRel}
}

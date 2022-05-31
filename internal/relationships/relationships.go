package relationships

import (
	"bytes"
	"strings"
)

type Relation string
type Relations []Relation

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
	BookRelation    = Relation("books")
)

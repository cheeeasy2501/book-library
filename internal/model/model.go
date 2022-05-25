package model

import (
	"encoding/json"
	"time"
)

type TestMapper interface {
	Columns() string
	Fields() []interface{}
}

type Model struct {
	Id        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (m *Model) ToMap() (map[string]interface{}, error) {
	var mp map[string]interface{}
	buf, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &m)
	if err != nil {
		return nil, err
	}

	return mp, err
}

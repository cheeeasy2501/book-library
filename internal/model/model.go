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
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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

package model

import (
	"encoding/json"
	"time"
)

type Model struct {
	CreatedAt time.Time `json:"createdAt" binding:"datetime"`
	UpdatedAt time.Time `json:"updatedAt" binding:"datetime"`
}

func (model *Model) ToMap() (map[string]interface{}, error) {
	var m map[string]interface{}
	buf, err := json.Marshal(model)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &m)
	if err != nil {
		return nil, err
	}

	return m, err
}

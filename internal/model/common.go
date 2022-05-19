package model

import (
	"encoding/json"
	"time"
)

type Model struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (b *Model) ToMap() (map[string]interface{}, error) {
	var m map[string]interface{}
	buf, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(buf, &m)
	if err != nil {
		return nil, err
	}

	return m, err
}

type PaginationParams struct {
	Page  uint64 `form:"page" json:"page" binding:"required,gte=1"`
	Limit uint64 `form:"limit" json:"limit" binding:"required,gte=1"`
}

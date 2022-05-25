package model

import (
	"time"
)

type TestMapper interface {
	Columns() string
	Fields() []interface{}
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

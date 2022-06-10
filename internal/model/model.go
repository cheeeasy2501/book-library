package model

import (
	"time"
)

type ScanMapper interface {
	ScanFields() []interface{}
}

type FieldSetter interface {
	Columns() string
}

type Timestamp struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

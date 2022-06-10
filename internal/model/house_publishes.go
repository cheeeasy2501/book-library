package model

import (
	"encoding/json"
	"errors"
)

type BookHousePublishes struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Timestamp
}

func (b *BookHousePublishes) Scan(src interface{}) error {
	bts, ok := src.([]byte)
	if !ok {
		return errors.New("Error Scanning Array")
	}

	return json.Unmarshal(bts, b)
}

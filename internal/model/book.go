package model

import (
	"encoding/json"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"time"
)

type Book struct {
	Id          uint64    `json:"id"`
	AuthorId    *int64    `json:"authorId"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	InStock     uint      `json:"inStock"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return apperrors.ValidateError("Invalid Title argument")
	}

	return nil
}

func (b *Book) ToMap() (map[string]interface{}, error) {
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

func (b *Book) UpdateMap() (map[string]interface{}, error) {
	m, err := b.ToMap()
	if err != nil {
		return nil, err
	}
	delete(m, "id")
	delete(m, "createdAt")

	return m, err
}

type GetBookParams struct {
	Id uint64 `uri:"id" binding:"required"`
}

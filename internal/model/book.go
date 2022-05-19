package model

import (
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
)

type Book struct {
	Model
	Id          uint64 `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Link        string `json:"link"`
	InStock     uint   `json:"inStock"`
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return apperrors.ValidateError("Invalid Title argument")
	}

	return nil
}

//type Books struct {
//	books[]
//}

type GetBookParams struct {
	Id uint64 `uri:"id" binding:"required"`
}

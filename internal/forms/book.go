package forms

import (
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/relationships"
	"github.com/gin-gonic/gin"
)

type GetBooksForm struct {
	Pagination
	relationships.Relations
}

func (f *GetBooksForm) LoadAndValidate(ctx *gin.Context) error {
	var err error

	err = f.Pagination.LoadAndValidate(ctx)
	if err != nil {
		return apperrors.ValidateError(err.Error())
	}
	err = f.Relations.LoadAndValidate(ctx)
	if err != nil {
		return apperrors.ValidateError(err.Error())
	}

	return nil
}
func NewGetBooksForm() *GetBooksForm {
	return &GetBooksForm{}
}

type GetBookForm struct {
	Id uint64 `uri:"id" binding:"required,gte=1"`
	relationships.Relations
}

func NewGetBookForm() *GetBookForm {
	return &GetBookForm{}
}

type CreateBookForm struct {
	HousePublishId uint64   `json:"house_publish_id" binding:"gte=0"`
	Title          string   `json:"title" binding:"required"`
	Description    string   `json:"description"`
	Link           string   `json:"link" binding:"url"`
	InStock        uint     `json:"in_stock" binding:"gte=0"`
	AuthorIds      []uint64 `json:"author_ids"`
}

func NewCreateBookForm() *CreateBookForm {
	return &CreateBookForm{}
}

type UpdateBookForm struct {
	Id             uint64 `uri:"id"`
	HousePublishId int64  `json:"house_publish_id" binding:"gte=0"`
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description"`
	Link           string `json:"link" binding:"url"`
	InStock        int    `json:"in_stock" binding:"qte=1"`
	AuthorIds      []int  `json:"author_ids"`
}

func NewUpdateBookForm() *UpdateBookForm {
	return &UpdateBookForm{}
}

type DeleteBookForm struct {
	Id uint64 `uri:"id"`
}

func NewDeleteBookForm() *DeleteBookForm {
	return &DeleteBookForm{}
}

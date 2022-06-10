package forms

import (
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/gin-gonic/gin"
)

type FormInterface interface {
	LoadAndValidate(ctx *gin.Context) error
}

type Pagination struct {
	Page   uint64 `form:"page" json:"page" binding:"required,gte=1"`
	Limit  uint64 `form:"limit" json:"limit" binding:"required,gte=1"`
	Offset uint64
}

func (pf *Pagination) LoadAndValidate(ctx *gin.Context) error {
	err := ctx.BindQuery(&pf)
	if err != nil {
		return apperrors.ValidateError(err.Error())
	}
	pf.Offset = (pf.Page - 1) * pf.Limit

	return nil
}

func (pf Pagination) GetOffset() uint64 {
	return (pf.Page - 1) * pf.Limit
}

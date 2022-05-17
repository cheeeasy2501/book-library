package app

import (
	"github.com/cheeeasy2501/book-library/internal/book"
	"github.com/cheeeasy2501/book-library/internal/errors"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/gin-gonic/gin"
)

func (a *App) GetBooks(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()
	params := model.GetBooksParams{}
	err = ctx.ShouldBindQuery(&params)
	if err != nil {
		err = errors.ValidateError(err.Error())
		return
	}
	books, err := a.service.Book.GetAll(ctx, params)
	if err != nil {
		return
	}

	a.SendResponse(ctx, books)
}

func (a *App) GetBook(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	uriParams := book.GetBookQuery{}
	err = ctx.ShouldBindUri(&uriParams)

	bk, err := a.bookController.GetBook(ctx, uriParams.Id)
	if err != nil {
		return
	}

	a.SendResponse(ctx, bk)
}

func (a *App) CreateBook(ctx *gin.Context) {

}

func (a *App) UpdateBook(ctx *gin.Context) {

}

func (a *App) DeleteBook(ctx *gin.Context) {

}

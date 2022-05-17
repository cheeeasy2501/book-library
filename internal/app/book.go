package app

import (
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
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
		err = apperrors.ValidateError(err.Error())
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

	uriParams := model.GetBookParams{}
	err = ctx.ShouldBindUri(&uriParams)
	bk, err := a.service.Book.GetById(ctx, uriParams.Id)
	if err != nil {
		return
	}

	a.SendResponse(ctx, bk)
}

func (a *App) CreateBook(ctx *gin.Context) {
	var (
		err  error
		book *model.Book
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.BindJSON(&book)
	if err != nil {
		return
	}

	err = a.service.Book.Create(ctx, book)
	if err != nil {
		return
	}

	a.SendResponse(ctx, book)
}

func (a *App) UpdateBook(ctx *gin.Context) {

}

func (a *App) DeleteBook(ctx *gin.Context) {

}

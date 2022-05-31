package app

import (
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/relationships"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (a *App) GetBooks(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	paginateForm := forms.Pagination{}
	err = ctx.BindQuery(&paginateForm)
	if err != nil {
		err = apperrors.ValidateError(err.Error())
		return
	}

	relations := relationships.Relations{}
	relationsQuery, ok := ctx.GetQuery("relations")
	if ok {
		err = relations.UnmarshalText([]byte(relationsQuery))
		if err != nil {
			return
		}
	}

	switch len(relations) {
	case 0:
		var books []model.Book
		books, err = a.service.Book.GetAll(ctx, paginateForm)
		if err != nil {
			return
		}
		a.SendResponse(ctx, books)
	default:
		var books []model.BookAggregate
		books, err = a.service.BookAggregate.GetAll(ctx, paginateForm, relations)
		if err != nil {
			return
		}
		a.SendResponse(ctx, books)
	}
}

func (a *App) GetBook(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	form := forms.GetBook{}
	err = ctx.BindUri(&form)

	relations := relationships.Relations{}
	relationsQuery, ok := ctx.GetQuery("relations")
	if ok {
		err = relations.UnmarshalText([]byte(relationsQuery))
		if err != nil {
			return
		}
	}
	count := len(relations)

	switch count {
	case 0:
		var book *model.Book
		book, err = a.service.Book.GetById(ctx, form.Id)
		a.SendResponse(ctx, book)
		if err != nil {
			return
		}
	default:
		var book *model.BookAggregate
		book, err = a.service.BookAggregate.GetById(ctx, form.Id, relations)
		if err != nil {
			return
		}
		a.SendResponse(ctx, book)
	}
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
	var (
		err  error
		book *model.Book
		id   uint64
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.BindJSON(&book)
	if err != nil {
		return
	}
	id, err = strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return
	}
	book.Id = id
	book.UpdatedAt = time.Now()

	err = a.service.Book.Update(ctx, book)
	if err != nil {
		return
	}

	a.SendResponse(ctx, book)
}

func (a *App) DeleteBook(ctx *gin.Context) {
	var (
		err error
		id  uint64
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	id, err = strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return
	}

	err = a.service.Book.Delete(ctx, id)
	if err != nil {
		return
	}

	a.SendResponse(ctx, nil)
}

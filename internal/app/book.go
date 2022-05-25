package app

import (
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
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

	rel := forms.Relationships{}
	// TODO: Problem with bind relations
	relationsQuery, ok := ctx.GetQuery("relations")
	if ok {
		err = rel.UnmarshalText([]byte(relationsQuery))
		if err != nil {
			return
		}
	}

	relations := rel.Relations.FilterRelations(forms.GetBookRelations())
	// Mock!
	//relations = forms.GetBookRelations()
	if len(relations) != 0 {
		books, err := a.service.BookAggregate.GetAll(ctx, paginateForm, relations)
		if err != nil {
			return
		}
		a.SendResponse(ctx, books)
	} else {
		books, err := a.service.Book.GetAll(ctx, paginateForm, relations)
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
	bk, err := a.service.Book.GetById(ctx, form.Id)
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

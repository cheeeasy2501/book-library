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

// GetBooks godoc
// @Summary      Get book collection
// @Description  Return book collection
// @Tags         books
// @Accept       json
// @Consume 	 json
// @Param        page    query  int  true  "Page number"
// @Param        limit    query  int  true  "Limit number"
// @Param        relations    query  string  false  "Book relationships:publish_house,author"
// @Success      200  {array}   model.Book
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /books [get]
func (a *App) GetBooks(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	pagination := forms.Pagination{}
	err = ctx.BindQuery(&pagination)
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

	var books []model.Book
	books, err = a.service.Book.GetAll(ctx, pagination, relations)
	if err != nil {
		return
	}
	a.SendResponse(ctx, books)

}

// GetBook godoc
// @Summary      Get book by id
// @Description  Return book
// @Tags         books
// @Accept       json
// @Consume 	 json
// @Param        id    path  int  true  "Book ID"
// @Param        relations    query  string  false  "Book relationships:authors,publish_house"
// @Success      200  {object}  model.Book
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /books/{id} [get]
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

	var book *model.Book
	book, err = a.service.Book.GetById(ctx, form.Id, relations)
	if err != nil {
		return
	}
	a.SendResponse(ctx, book)
}

// GetBook godoc
// @Summary      Registration new account
// @Description  Return user and token
// @Tags         books
// @Accept       json
// @Consume 	 json
// @Param        page    query  string  true  "Page number"
// @Param        limit    query  string  true  "Limit number"
// @Param        relations    query  string  true  "Book relationships:authors,publish_house"
// @Success      200  {object}  model.Book
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
func (a *App) CreateBook(ctx *gin.Context) {
	var (
		err  error
		book *model.Book
		//authors *model.Authors
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

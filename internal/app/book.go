package app

import (
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/gin-gonic/gin"
	"time"
)

func (a *App) GetBooks(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	form := forms.NewGetBooksForm()
	err = form.LoadAndValidate(ctx)
	if err != nil {
		return
	}
	_, ok := ctx.GetQuery("relations")

	switch ok {
	case true:
		book, err := a.service.Book.GetAllWithRelations(ctx, form.Pagination)
		if err != nil {
			return
		}
		a.SendResponse(ctx, book)
	default:
		books, err := a.service.Book.GetAll(ctx, form.Pagination)
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

	form := forms.NewGetBookForm()
	err = ctx.ShouldBindUri(&form)
	if err != nil {
		return
	}
	_, ok := ctx.GetQuery("relations")

	switch ok {
	case true:
		book, err := a.service.Book.GetByIdWithRelations(ctx, form.Id)
		if err != nil {
			return
		}
		a.SendResponse(ctx, book)
	default:
		book, err := a.service.Book.GetById(ctx, form.Id)
		if err != nil {
			return
		}
		a.SendResponse(ctx, book)
	}
}

func (a *App) CreateBook(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	form := forms.NewCreateBookForm()
	err = ctx.BindJSON(&form)
	if err != nil {
		return
	}

	book := model.CreateBook{
		Book: model.Book{
			HousePublishId: form.HousePublishId,
			Title:          form.Title,
			Description:    form.Description,
			Link:           form.Link,
			InStock:        form.InStock,
		},
		AuthorIds: form.AuthorIds,
	}
	fullBook, err := a.service.Book.CreateWithRelations(ctx, book)
	if err != nil {
		return
	}
	a.SendResponse(ctx, fullBook)

}

func (a *App) UpdateBook(ctx *gin.Context) {
	var (
		err  error
		book *model.Book
	)

	defer func() {
		a.SendError(ctx, err)
	}()
	form := forms.NewUpdateBookForm()
	err = ctx.BindJSON(&form)
	if err != nil {
		return
	}
	err = ctx.BindUri(&form)
	if err != nil {
		return
	}
	book.UpdatedAt = time.Now()

	err = a.service.Book.Update(ctx, book)
	if err != nil {
		return
	}

	a.SendResponse(ctx, book)
}

// DeleteBook godoc
// @Summary      Delete book
// @Description  Delete book by id
// @Tags         books
// @Accept       json
// @Consume 	 json
// @Param   	 request  path    int     true  "FullBook id"
// @Success      200  {object}  model.FullBook
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /books/{id} [delete]
func (a *App) DeleteBook(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()
	form := forms.NewDeleteBookForm()
	err = ctx.ShouldBindUri(form)
	if err != nil {
		return
	}

	err = a.service.Book.Delete(ctx, form.Id)
	if err != nil {
		return
	}

	a.SendResponse(ctx, nil)
}

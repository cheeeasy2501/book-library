package app

import (
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/gin-gonic/gin"
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
		err   error
		books []model.Book
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	form := forms.NewGetBooksForm()
	err = form.LoadAndValidate(ctx)
	if err != nil {
		return
	}

	books, err = a.service.Book.GetAll(ctx, form.Pagination, form.Relations)
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
		err  error
		book *model.Book
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	form := forms.NewGetBookForm()
	err = ctx.ShouldBindUri(&form)
	if err != nil {
		return
	}

	relationsQuery, ok := ctx.GetQuery("relations")
	if ok {
		err = form.Relations.UnmarshalText([]byte(relationsQuery))
		if err != nil {
			return
		}
	}

	book, err = a.service.Book.GetById(ctx, form.Id, form.Relations)
	if err != nil {
		return
	}
	a.SendResponse(ctx, book)
}

// CreateBook godoc
// @Summary      Create new book
// @Description  Create new book and return it
// @Tags         books
// @Accept       json
// @Consume 	 json
// @Param   	 request  body  model.Book     true  "Create book model"
// @Success      200  {object}  model.Book
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /books [post]
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

	book := &model.Book{
		HousePublishId: form.HousePublishId,
		Title:          form.Title,
		Description:    form.Description,
		Link:           form.Link,
		InStock:        form.InStock,
	}
	err = a.service.Book.Create(ctx, book)
	if err != nil {
		return
	}

	a.SendResponse(ctx, book)
}

// UpdateBook godoc
// @Summary      Update book
// @Description  Update book and return it
// @Tags         books
// @Accept       json
// @Consume 	 json
// @Param   	 id  path  int     true  "Book id"
// @Param   	 request  body  model.Book     true  "Update book model"
// @Success      200  {object}  model.Book
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /books/{id} [patch]
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
// @Param   	 request  path    int     true  "Book id"
// @Success      200  {object}  model.Book
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

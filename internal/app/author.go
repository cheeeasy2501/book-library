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

// Get authors godoc
// @Summary      Get author collection
// @Description  Return author collection
// @Tags         author
// @Accept       json
// @Consume 	 json
// @Param        page    query  int  true  "Page number"
// @Param        limit    query  int  true  "Limit number"
// @Param        relations    query  string  false  "Author relationships:books"
// @Success      200  {array}   model.Author
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /authors [get]
func (a *App) GetAuthors(ctx *gin.Context) {
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

	var authors []model.Author
	authors, err = a.service.Author.GetAll(ctx, pagination, relations)
	if err != nil {
		return
	}
	a.SendResponse(ctx, authors)
}

// GetAuthor godoc
// @Summary      Get author by id
// @Description  Return author
// @Tags         author
// @Accept       json
// @Consume 	 json
// @Param        id    path  int  true  "Author id"
// @Param        relations    query  string  false  "Author relationships:books"
// @Success      200  {object}  model.Author
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /authors/{id} [get]
func (a *App) GetAuthor(ctx *gin.Context) {
	var (
		err error
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	form := forms.GetAuthorForm{}
	err = ctx.BindUri(&form)

	relations := relationships.Relations{}
	relationsQuery, ok := ctx.GetQuery("relations")
	if ok {
		err = relations.UnmarshalText([]byte(relationsQuery))
		if err != nil {
			return
		}
	}

	var author *model.Author
	author, err = a.service.Author.GetById(ctx, form.Id, relations)
	if err != nil {
		return
	}
	a.SendResponse(ctx, author)
}

// CreateAuthor godoc
// @Summary      Create new author
// @Description  Create new author and return it
// @Tags         author
// @Accept       json
// @Consume 	 json
// @Param   	 request  body  model.Author     true  "Create author model"
// @Success      200  {object}  model.Author
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /authors [post]
func (a *App) CreateAuthor(ctx *gin.Context) {
	var (
		err    error
		author *model.Author
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.BindJSON(&author)
	if err != nil {
		return
	}

	err = a.service.Author.Create(ctx, author)
	if err != nil {
		return
	}

	a.SendResponse(ctx, author)
}

// UpdateAuthor godoc
// @Summary      Update author
// @Description  Update author and return it
// @Tags         author
// @Accept       json
// @Consume 	 json
// @Param   	 id  path  int     true  "author id"
// @Param   	 request  body  model.Author     true  "Update author model"
// @Success      200  {object}  model.Author
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Security     ApiKeyAuth
// @Router       /authors/{id} [patch]
func (a *App) UpdateAuthor(ctx *gin.Context) {
	var (
		err    error
		author *model.Author
		id     uint64
	)

	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.BindJSON(&author)
	if err != nil {
		return
	}
	id, err = strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return
	}
	author.Id = id
	author.UpdatedAt = time.Now()

	err = a.service.Author.Update(ctx, author)
	if err != nil {
		return
	}

	a.SendResponse(ctx, author)
}

// Deleteauthor godoc
// @Summary      Delete author
// @Description  Delete author by id
// @Tags         author
// @Accept       json
// @Consume 	 json
// @Param   	 request  path    int     true  "author id"
// @Success      200  {object}  model.Author
// @Failure      400  {object}  HTTPError
// @Failure      404  {object}  HTTPError
// @Failure      500  {object}  HTTPError
// @Router       /authors/{id} [delete]
func (a *App) DeleteAuthor(ctx *gin.Context) {
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

	err = a.service.Author.Delete(ctx, id)
	if err != nil {
		return
	}

	a.SendResponse(ctx, nil)
}

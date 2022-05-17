package model

import "time"

type Book struct {
	ID          uint64    `json:"id"`
	AuthorID    *int64    `json:"authorId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	InStock     uint      `json:"inStock"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type GetBooksParams struct {
	Page  uint64 `form:"page" json:"page" binding:"required,gte=1"`
	Limit uint64 `form:"limit" json:"limit" binding:"required,gte=1"`
}

type GetBookQuery struct {
	Id uint64 `uri:"id" binding:"required"`
}

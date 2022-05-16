package book

import "github.com/google/uuid"

type Book struct {
	ID          uuid.UUID `json:"id"`
	AuthorID    int64     `json:"authorId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

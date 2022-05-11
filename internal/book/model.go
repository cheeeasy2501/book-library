package book

import "github.com/google/uuid"

type Book struct {
	ID          uuid.UUID `json:"ID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
}

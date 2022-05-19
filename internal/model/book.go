package model

type Book struct {
	Model
	Id          uint64 `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Link        string `json:"link" binding:"url"`
	InStock     uint   `json:"inStock"`
}

type BookAggregate struct {
	Book
	Authors []Author `json:"authors,omitempty"`
}

package model

type Model interface{}

type Book struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Link        string `json:"link" binding:"url"`
	InStock     uint   `json:"in_stock"`
	Timestamp
}

package model

type Book struct {
	Id             uint64 `json:"id"`
	HousePublishId uint64 `json:"house_publish_id"`
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description"`
	Link           string `json:"link" binding:"url"`
	InStock        uint   `json:"in_stock"`
	Timestamp
}

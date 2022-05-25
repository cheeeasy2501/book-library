package model

type Book struct {
	Model
	PublishHouseId uint64 `json:"publish_house_id"`
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description"`
	Link           string `json:"link" binding:"url"`
	InStock        uint   `json:"in_stock"`
}

func (a *Book) Columns() string {
	return "author.id, author.firstname, author.lastname, author.created_at, author.updated_at"
}

func (a *Book) Fields() []interface{} {
	return []interface{}{&a.Id, &a.Title, &a.Description, &a.Link, &a.InStock, &a.CreatedAt, a.UpdatedAt}
}

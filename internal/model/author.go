package model

type Author struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Timestamp
	Books []Book `json:"books,omitempty"`
}

func (a *Author) Columns() string {
	return "author.id, author.firstname, author.lastname, author.created_at, author.updated_at"
}

func (a *Author) Fields() []interface{} {
	return []interface{}{&a.Id, &a.FirstName, &a.LastName, &a.CreatedAt, a.UpdatedAt}
}

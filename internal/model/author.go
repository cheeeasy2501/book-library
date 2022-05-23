package model

type Author struct {
	Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Books     []Book `json:"books,omitempty"`
}

func (a *Author) Columns() string {
	return "author.id, author.firstname, author.lastname, author.created_at, author.updated_at"
}

func (a *Author) Fields() []interface{} {
	return []interface{}{&a.Id, &a.FirstName, &a.LastName, &a.CreatedAt, a.UpdatedAt}
}

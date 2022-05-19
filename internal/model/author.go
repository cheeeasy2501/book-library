package model

type Author struct {
	Model
	Id        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Books     []Book `json:"books"`
}

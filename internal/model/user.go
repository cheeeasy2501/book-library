package model

type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"userName" validate:"required"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type SignIn struct {
}

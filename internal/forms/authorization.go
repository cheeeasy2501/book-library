package forms

import "github.com/cheeeasy2501/book-library/internal/model"

type Credentials struct {
	UserName string `json:"username" example:"administrator"`
	Password string `json:"password" example:"administrator"`
}

type CreateUser struct {
	UserName  string  `json:"username" example:"Username"`
	Password  string  `json:"password" example:"Password"`
	FirstName *string `json:"firstname,omitempty" example:"MyFirstName"`
	LastName  *string `json:"lastname,omitempty" example:"MyLastName"`
	Email     *string `json:"email,omitempty" example:"MyEmail@example.com"`
}

func (cu *CreateUser) ToUserModel(user *model.User) {
	user.UserName = cu.UserName
	user.SetPassword(cu.Password)
	user.FirstName = cu.FirstName
	user.LastName = cu.LastName
	user.Email = cu.Email
}

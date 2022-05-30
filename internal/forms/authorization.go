package forms

import "github.com/cheeeasy2501/book-library/internal/model"

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type CreateUser struct {
	UserName  string  `json:"username"`
	Password  string  `json:"password"`
	FirstName *string `json:"firstname,omitempty"`
	LastName  *string `json:"lastname,omitempty"`
	Email     *string `json:"email,omitempty"`
}

func (cu *CreateUser) ToUserModel(user *model.User) {
	user.UserName = cu.UserName
	user.SetPassword(cu.Password)
	user.FirstName = cu.FirstName
	user.LastName = cu.LastName
	user.Email = cu.Email
}

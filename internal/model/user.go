package model

type User struct {
	Model
	Id        int64  `json:"id"`
	UserName  string `json:"userName" validate:"required"`
	password  string
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email" validate:"email"`
}

func (user *User) Password() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = password
}

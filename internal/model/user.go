package model

type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"userName" validate:"required"`
	password  string
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email" validate:"email"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

func (user *User) Password() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = password
}

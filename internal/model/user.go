package model

type User struct {
	Model
	Id        int64  `json:"id"`
	UserName  string `json:"userName"`
	password  string
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email"`
}

func (user *User) Password() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = password
}

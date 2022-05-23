package model

type User struct {
	Model
	Id        int64  `json:"id"`
	UserName  string `json:"userName"`
	password  string
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Email     *string `json:"email,omitempty"`
}

func (user *User) Password() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = password
}

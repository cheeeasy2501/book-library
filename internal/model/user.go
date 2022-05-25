package model

type User struct {
	Model
	Id        int64  `json:"id"`
	UserName  string `json:"username"`
	password  string
	FirstName *string `json:"firstname,omitempty"`
	LastName  *string `json:"lastname,omitempty"`
	Email     *string `json:"email,omitempty"`
}

func (user *User) Password() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = password
}

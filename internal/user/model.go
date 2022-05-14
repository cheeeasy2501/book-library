package user

type User struct {
	Id        int64  `json:"id"`
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

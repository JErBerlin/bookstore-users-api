package users

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func NewUser(id int64, firstName string, lastName string, email string, dateCreated string) *User {
	return &User{Id: id, FirstName: firstName, LastName: lastName, Email: email, DateCreated: dateCreated}
}

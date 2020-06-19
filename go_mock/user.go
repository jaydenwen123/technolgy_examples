package user

//User
type User struct {
	Name string
}

type UserRepository interface {
	FindOne(id int) (*User, error)
}

//NewUser
func NewUser() *User {
	return &User{}
}


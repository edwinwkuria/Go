package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	LastId = 0
)

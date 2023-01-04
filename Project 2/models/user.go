package models

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	NextId = 0
)

func GetUser() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.Id = NextId
	NextId++
	users = append(users, &u)
	return u, nil
}

package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // slice of pointers of users
	nextID = 1
)

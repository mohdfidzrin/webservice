package models

import "time"

type User struct {
	ID int
	Name string
	DOB time.Time
}

var (
	users []*User
	nextID = 1
)

func FindAll() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}


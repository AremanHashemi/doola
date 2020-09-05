package user

import (
	"time"
)

type FakeRepo struct {
	users []User
}

func (r FakeRepo) CreateUser(name string) User {
	u := User{
		Id: len(r.users),
		CreatedAt: time.Now(),
		Name: name,
	}

	r.users = append(r.users, u)

	return u
}

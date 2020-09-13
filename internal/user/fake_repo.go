package user

import (
	"errors"
	"fmt"
	"time"
)

type FakeRepo struct {
	users []User
}

func (r *FakeRepo) CreateUser(name string) User {
	u := User{
		Id:        len(r.users),
		CreatedAt: time.Now(),
		Name:      name,
	}
	r.users = append(r.users, u)
	return u
}

func (r *FakeRepo) GetUser(id int) (User, error) {
	fmt.Print()
	for _, u := range r.users {
		if u.Id == id {
			return u, nil
		}
	}
	return User{}, errors.New("UserNotFound")
}

package doola

import "time"

type FakeRepo struct {
	users []User
}

func (r FakeRepo) CreateUser(name string) User {
	u := User{
		id: len(r.users),
		createdAt: time.Now(),
		name: name,
	}

	r.users = append(r.users, u)

	return u
}

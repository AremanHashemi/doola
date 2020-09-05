package internal

import "time"

type User struct {
	id int
	createdAt time.Time
	name string
}

type Repository interface {
	CreateUser(name string) User
}

type Controller struct {
	Repo Repository
}

func (c Controller) CreateUser(name string) User {
	return c.Repo.CreateUser(name)
}

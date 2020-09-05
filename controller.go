package doola

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
	repo Repository
}

func (c Controller) CreateUser(name string) User {
	return c.repo.CreateUser(name)
}

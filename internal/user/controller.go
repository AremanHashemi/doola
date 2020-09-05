package user

import "time"

type User struct {
	Id        int
	CreatedAt time.Time
	Name      string
}

type Repository interface {
	CreateUser(name string) User
}

type Controller struct {
	repo Repository
}

func NewController(r Repository) Controller {
	return Controller{
		repo: r,
	}
}

func (c Controller) CreateUser(name string) User {
	return c.repo.CreateUser(name)
}

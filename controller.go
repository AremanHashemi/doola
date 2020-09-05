package doola

type User struct {
	name string
}

type Repository interface {
	CreateUser(name string) User
}

type Controller struct {
	repo Repository
}

func (c Controller) CreateUser(name string) User {
	return c.CreateUser(name)
}

package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestController_CreateUser(t *testing.T) {
	r := FakeRepo{}
	c := Controller{
		repo: &r,
	}

	u := c.CreateUser("Areman")

	assert.Equal(t, "Areman", u.Name)

}

func TestController_GetUser(t *testing.T) {
	r := FakeRepo{}
	c := NewController(&r)
	u := c.CreateUser("Ben")

	a, err := c.GetUser(u.Id)
	if err != nil {
		assert.Fail(t, "error on getUser")
	}
	assert.Equal(t, "Ben", a.Name)
}

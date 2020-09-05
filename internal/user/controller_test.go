package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestController_CreateUser(t *testing.T) {
	r := FakeRepo{}
	c := Controller{
		repo: r,
	}

	u := c.CreateUser("Areman")

	assert.Equal(t, "Areman", u.Name)
}

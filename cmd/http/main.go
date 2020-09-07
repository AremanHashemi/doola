package main

import (
	"github.com/camhashemi/doola/internal/server"
	"github.com/camhashemi/doola/internal/user"
)

func main() {
	r := user.FakeRepo{}
	c := user.NewController(r)
	s := server.Server{
		Controller: c,
	}

	s.Start()
}

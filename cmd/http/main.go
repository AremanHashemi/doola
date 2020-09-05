package main

import (
	"doola/internal/server"
	"doola/internal/user"
)

func main() {
	r := user.FakeRepo{}
	c := user.NewController(r)
	s := server.Server{
		Controller: c,
	}

	s.Start()
}

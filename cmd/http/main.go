package main

import (
	"doola/internal/http"
	"doola/internal/user"
)

func main() {
	r := user.FakeRepo{}
	c := user.NewController(r)
	s := http.Server{
		Controller: c,
	}

	s.Start()
}

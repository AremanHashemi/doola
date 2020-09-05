package main

import "doola/internal"

func main() {
	r := internal.FakeRepo{}
	c := internal.Controller{
		Repo: r,
	}
	s := internal.Server{
		Controller: c,
	}

	s.Start()
}

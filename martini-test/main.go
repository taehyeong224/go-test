package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello, 세계!"
	})
	m.Run()
}
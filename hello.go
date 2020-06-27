package main

import (
	"./person"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.GET("/person", person.Index)
	router.GET("/person/hello/:name", person.Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}

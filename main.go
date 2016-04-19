package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/supergiant/universe/routectl"
)

func main() {
	router := httprouter.New()
	router.GET("/", routectl.Index)
	router.GET("/search/*component", routectl.Search)
	router.GET("/example", routectl.Example)

	log.Println("Universe API Running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

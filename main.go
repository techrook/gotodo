package main

import (
	"log"
	"net/http"

	"github.com/techrook/gotodo/routes"
)

func main() {
	err := http.ListenAndServe(":3333", routes.Init())

	if err != nil {
		log.Fatal(err)
	}
}

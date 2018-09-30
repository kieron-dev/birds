package main

import (
	"net/http"

	"github.com/kieron-pivotal/birdpedia/birds/handlers"
	"github.com/kieron-pivotal/birdpedia/birds/storage"
	"github.com/kieron-pivotal/birdpedia/routes"
)

func main() {
	s := new(storage.Birds)
	birdHandler := handlers.NewHandler(s)
	r := routes.NewRouter(birdHandler)
	http.ListenAndServe(":8080", r)
}

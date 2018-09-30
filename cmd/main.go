package main

import (
	"net/http"

	"github.com/kieron-pivotal/birdpedia"
	"github.com/kieron-pivotal/birdpedia/birds/handlers"
	"github.com/kieron-pivotal/birdpedia/birds/storage"
)

func main() {
	s := new(storage.Birds)
	birdHandler := handlers.NewHandler(s)
	r := birdpedia.NewRouter(birdHandler)
	http.ListenAndServe(":8080", r)
}

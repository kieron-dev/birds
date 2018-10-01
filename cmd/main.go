package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/kieron-pivotal/birdpedia/birds/handlers"
	"github.com/kieron-pivotal/birdpedia/birds/storage/db"
	"github.com/kieron-pivotal/birdpedia/routes"
)

func main() {
	// s := new(memory.Store)
	connString := "host=/var/run/postgresql dbname=bird_encyclopedia sslmode=disable"
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	s := db.NewDBStore(conn)

	birdHandler := handlers.NewHandler(s)
	r := routes.NewRouter(birdHandler)
	http.ListenAndServe(":8080", r)
}

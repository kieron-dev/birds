//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/kieron-pivotal/birdpedia/birds/handlers"
	"github.com/kieron-pivotal/birdpedia/birds/storage"
	"github.com/kieron-pivotal/birdpedia/birds/storage/db"
	"github.com/kieron-pivotal/birdpedia/database"
	"github.com/kieron-pivotal/birdpedia/routes"
)

func InitialiseRouter() (*mux.Router, error) {
	wire.Build(
		routes.NewRouter,
		handlers.NewHandler,

		// using memory data store:
		// wire.Struct(new(memory.Store)),
		// wire.Bind(new(storage.Store), new(*memory.Store)),

		// using postgres data store:
		wire.Value(database.ConnectionString("host=/var/run/postgresql dbname=bird_encyclopedia sslmode=disable")),
		database.NewPostgresConnection,
		db.NewDBStore,
		wire.Bind(new(storage.Store), new(*db.Store)),
	)
	return &mux.Router{}, nil
}

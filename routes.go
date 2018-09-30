package birdpedia

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kieron-pivotal/birdpedia/birds/handlers"
	"github.com/kieron-pivotal/birdpedia/hello"
)

func NewRouter(birdHandler handlers.Handler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", hello.Handler).Methods("GET")

	assetsDir := fmt.Sprintf("%s/src/github.com/kieron-pivotal/birdpedia/assets/", os.Getenv("GOPATH"))
	staticFileDir := http.Dir(assetsDir)
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/bird", birdHandler.GetBirdsHandler).Methods("GET")
	r.HandleFunc("/bird", birdHandler.CreateBirdHandler).Methods("POST")
	return r
}

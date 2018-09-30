package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kieron-pivotal/birdpedia/birds"
	"github.com/kieron-pivotal/birdpedia/birds/storage"
)

type Handler struct {
	storage *storage.Birds
}

func NewHandler(birdStorage *storage.Birds) Handler {
	return Handler{
		storage: birdStorage,
	}
}

func (h Handler) GetBirds(w http.ResponseWriter, r *http.Request) {
	birdJson, err := json.Marshal(h.storage.GetList())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(birdJson)
}

func (h Handler) CreateBird(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird := birds.Bird{
		Species:     r.Form.Get("species"),
		Description: r.Form.Get("description"),
	}

	h.storage.Add(bird)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}

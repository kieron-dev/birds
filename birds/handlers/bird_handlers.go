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
	storage storage.Store
}

func NewHandler(birdStorage storage.Store) Handler {
	return Handler{
		storage: birdStorage,
	}
}

func (h Handler) GetBirds(w http.ResponseWriter, r *http.Request) {
	birdList, err := h.storage.GetBirds()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	birdJson, err := json.Marshal(birdList)
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

	err = h.storage.CreateBird(&bird)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}

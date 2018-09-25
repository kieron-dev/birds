package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdsHandler(w http.ResponseWriter, r *http.Request) {
	birdJson, err := json.Marshal(birds)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(birdJson)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird := Bird{
		Species:     r.Form.Get("species"),
		Description: r.Form.Get("description"),
	}
	fmt.Printf("bird = %+v\n", bird)
	birds = append(birds, bird)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}

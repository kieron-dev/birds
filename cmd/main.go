package main

import (
	"net/http"
)

func main() {
	r, err := InitialiseRouter()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", r)
}

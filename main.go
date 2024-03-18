package main

import (
	"fmt"
	"net/http"
	"scorecard-server/repository"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	repository.AddScore(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user", addUser)
	http.ListenAndServe(":8080", nil)
}

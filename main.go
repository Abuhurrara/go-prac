package main

import (
	"log"
	"net/http"
)

type api struct {
	address string
}

func (s *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User List..."))
}

func (s *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created User"))
}

func main() {
	api := &api{address: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.address,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUserHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

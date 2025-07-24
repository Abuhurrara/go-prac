package main

import (
	"log"
	"net/http"
)

type server struct {
	address string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server!"))
}

func main() {
	s := &server{address: ":8080"}
	if err := http.ListenAndServe(s.address, s); err != nil {
		log.Fatal(err)
	}
}

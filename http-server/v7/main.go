package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(&InMemoryPlayerStore{})
	log.Fatal(http.ListenAndServe(":8080", server))
}

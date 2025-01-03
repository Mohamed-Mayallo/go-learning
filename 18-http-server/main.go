package main

import (
	"log"
	"net/http"
)

func main() {
	db := InMemoryDb{store: map[string]int{}}
	mux := InitPlayerServer(&db)

	log.Fatal(http.ListenAndServe(":5000", mux))
}

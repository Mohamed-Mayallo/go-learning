package main

import (
	"log"
	"net/http"
)

func main() {
	db := InMemoryDb{store: map[string]int{}}

	server := PlayerServer{}
	mux := server.InitPlayerServer(&db)

	log.Fatal(http.ListenAndServe(":5000", mux))
}

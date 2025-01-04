package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	store, err := os.OpenFile("db.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening db.json %v", err)
	}

	db := FileSystemPlayerStore{store}
	mux := InitPlayerServer(&db)

	log.Fatal(http.ListenAndServe(":5000", mux))
}

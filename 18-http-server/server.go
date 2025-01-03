package main

import (
	"fmt"
	"net/http"
	"sync"
)

type PlayerServer struct {
	mux *http.ServeMux
}

var mu sync.RWMutex

func (s *PlayerServer) InitPlayerServer(db *InMemoryDb) *http.ServeMux {
	s.mux = http.NewServeMux()

	s.mux.HandleFunc("GET /players/{name}", GetScore(db))
	s.mux.HandleFunc("POST /players/{name}", CreateScore(db))

	return s.mux
}

func GetScore(db *InMemoryDb) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		playerName := r.PathValue("name")
		mu.Lock()
		score, ok := db.Get(playerName)
		mu.Unlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprint(w, score)
		// w.WriteHeader(http.StatusOK) // No need for this as http.ResponseWriter.Write implicitly calls w.WriteHeader(http.StatusOK)
	}
}

func CreateScore(db *InMemoryDb) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		playerName := r.PathValue("name")
		mu.Lock()
		db.Add(playerName)
		mu.Unlock()
		w.WriteHeader(http.StatusAccepted)
	}
}

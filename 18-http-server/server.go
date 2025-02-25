package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type PlayerServer struct {
	db *FileSystemPlayerStore
	http.Handler
}

type Player struct {
	Name  string
	Score int
}

var mu sync.RWMutex

func InitPlayerServer(db *FileSystemPlayerStore) PlayerServer {
	server := PlayerServer{db: db}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /players", GetPlayers(db))
	mux.HandleFunc("GET /players/{name}", GetScore(db))
	mux.HandleFunc("POST /players/{name}", CreateScore(db))

	server.Handler = mux

	return server
}

func GetScore(db *FileSystemPlayerStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		playerName := r.PathValue("name")
		mu.Lock()
		score, ok := db.GetPlayerScore(playerName)
		mu.Unlock()
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprint(w, score)
		// w.WriteHeader(http.StatusOK) // No need for this as http.ResponseWriter.Write implicitly calls w.WriteHeader(http.StatusOK)
	}
}

func CreateScore(db *FileSystemPlayerStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		playerName := r.PathValue("name")
		mu.Lock()
		db.Add(playerName)
		mu.Unlock()
		w.WriteHeader(http.StatusAccepted)
	}
}

func GetPlayers(db *FileSystemPlayerStore) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		players, _ := db.GetMany()
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(players)
	}
}

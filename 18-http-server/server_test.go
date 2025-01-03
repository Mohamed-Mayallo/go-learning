package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestGetPlayer(t *testing.T) {
	store := map[string]int{
		"M": 30,
		"A": 20,
	}
	db := InMemoryDb{store}
	mux := InitPlayerServer(&db)

	t.Run("player M", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/players/M", nil)
		res := httptest.NewRecorder()

		mux.ServeHTTP(res, req)

		got := res.Body.String()
		want := "30"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("player A", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/players/A", nil)
		res := httptest.NewRecorder()

		mux.ServeHTTP(res, req)

		got := res.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/players/S", nil)
		res := httptest.NewRecorder()

		mux.ServeHTTP(res, req)

		if res.Code != http.StatusNotFound {
			t.Error("S should be not found")
		}
	})
}

func TestPostPlayers(t *testing.T) {
	db := InMemoryDb{store: map[string]int{}}
	mux := InitPlayerServer(&db)

	t.Run("add player D", func(t *testing.T) {
		req1 := httptest.NewRequest(http.MethodPost, "/players/D", nil)
		res1 := httptest.NewRecorder()

		mux.ServeHTTP(res1, req1)

		req2 := httptest.NewRequest(http.MethodGet, "/players/D", nil)
		res2 := httptest.NewRecorder()

		mux.ServeHTTP(res2, req2)

		got := res2.Body.String()
		want := "1"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestGetPlayers(t *testing.T) {
	db := InMemoryDb{store: map[string]int{
		"A": 1,
		"B": 2,
	}}
	mux := InitPlayerServer(&db)

	t.Run("get many players", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/players", nil)
		res := httptest.NewRecorder()

		mux.ServeHTTP(res, req)

		var got []Player
		json.NewDecoder(res.Body).Decode(&got)

		want := []Player{
			{Name: "A", Score: 1},
			{Name: "B", Score: 2},
		}

		if !slices.Equal(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	store := map[string]int{
		"M": 30,
		"A": 20,
	}
	db := InMemoryDb{store}
	server := PlayerServer{}
	mux := server.InitPlayerServer(&db)

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
	server := PlayerServer{}
	mux := server.InitPlayerServer(&db)

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

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		fastServer := createMockServer(0 * time.Millisecond)
		slowServer := createMockServer(20 * time.Millisecond)

		defer fastServer.Close()
		defer slowServer.Close()

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		got, err := Racer(fastUrl, slowUrl)

		want := fastUrl

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := createMockServer(25 * time.Millisecond)

		defer server.Close()

		url := server.URL

		_, err := ConfigurableRacer(url, url, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expect an error here")
		}

	})
}

func createMockServer(sleepTime time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleepTime)
		w.WriteHeader(http.StatusOK)
	}))
}

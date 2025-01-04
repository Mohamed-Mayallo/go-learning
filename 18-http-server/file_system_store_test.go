package main

import (
	"os"
	"slices"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get players", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Score": 10},
			{"Name": "Chris", "Score": 33}]`)

		defer removeFile()

		store := FileSystemPlayerStore{database}

		got, err := store.GetMany()
		if err != nil {
			panic(err)
		}

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get players sorted by score", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "A", "Score": 10},
			{"Name": "G", "Score": 45},
			{"Name": "H", "Score": 33}]`)

		defer removeFile()

		store := FileSystemPlayerStore{database}

		got, err := store.GetMany()
		if err != nil {
			panic(err)
		}

		want := []Player{
			{"G", 45},
			{"H", 33},
			{"A", 10},
		}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("get player score", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Score": 10},
			{"Name": "Chris", "Score": 33}]`)

		defer removeFile()

		store := FileSystemPlayerStore{database}

		got, _ := store.GetPlayerScore("Chris")

		want := 33

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("add score to existing player", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Score": 10},
			{"Name": "Chris", "Score": 33}]`)

		defer removeFile()

		store := FileSystemPlayerStore{database}

		store.Add("Chris")

		got, _ := store.GetPlayerScore("Chris")
		want := 34

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("add score to non existing player", func(t *testing.T) {
		database, removeFile := createTempFile(t, `[
			{"Name": "Cleo", "Score": 10},
			{"Name": "Chris", "Score": 33}]`)

		defer removeFile()

		store := FileSystemPlayerStore{database}

		store.Add("MM")

		got, _ := store.GetPlayerScore("MM")
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func createTempFile(t testing.TB, initValue string) (*os.File, func()) {
	t.Helper()

	f, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	f.Write([]byte(initValue))

	removeFile := func() {
		f.Close()
		os.Remove(f.Name())
	}

	return f, removeFile
}

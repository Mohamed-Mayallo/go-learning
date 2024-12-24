package main

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		got, _ := dictionary.Search("test")

		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		_, err := dictionary.Search("another_term")

		if err == nil {
			t.Fatal("should have an error here")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new term", func(t *testing.T) {
		dic := Dictionary{}

		dic.Add("new_term", "val")

		got, _ := dic.Search("new_term")
		want := "val"

		assertStrings(t, got, want)
	})

	t.Run("existing term", func(t *testing.T) {
		dic := Dictionary{"test": "val"}

		err := dic.Add("test", "another val")

		assertError(t, err, ErrAlreadyFound)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dic := Dictionary{"test": "val"}

		dic.Update("test", "new val")

		got, _ := dic.Search("test")

		want := "new val"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dic := Dictionary{"test": "val"}

		err := dic.Update("another_test", "new val")

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dic := Dictionary{"test": "val"}

		dic.Delete("test")

		got, _ := dic.Search("test")

		want := ""

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dic := Dictionary{"test": "val"}

		err := dic.Delete("another_test")

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if !errors.Is(got, want) {
		t.Errorf("got error %q want %q", got, want)
	}
}

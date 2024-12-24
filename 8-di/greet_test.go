package main

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "hello")

	want := "hello"
	got := buffer.String()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

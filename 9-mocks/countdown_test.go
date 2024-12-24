package main

import (
	"bytes"
	"slices"
	"testing"
)

type SpySleeper struct {
	log []string
}

func (s *SpySleeper) Sleep() {
	s.log = append(s.log, "sleep")
}

func (s *SpySleeper) Write(p []byte) (n int, err error) {
	s.log = append(s.log, "write")
	return
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	spy := SpySleeper{log: []string{}}

	Countdown(&buffer, &spy)

	got := buffer.String()
	want := `3
2
1
Go!
`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	wantedLog := []string{
		"write", "sleep", "write", "sleep", "write", "sleep",
	}

	if slices.Equal(wantedLog, spy.log) {
		t.Errorf("log is out of order")
	}
}

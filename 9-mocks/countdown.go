package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct {
	duration time.Duration
	sleep    func(d time.Duration)
}

func (s DefaultSleeper) Sleep() {
	s.sleep(s.duration)
}

func main() {
	sleeper := DefaultSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}

type Sleeper interface {
	Sleep()
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprintln(w, "Go!")
}

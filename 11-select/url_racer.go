package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string) (string, error) {
	return ConfigurableRacer(url1, url2, 10*time.Second)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (string, error) {
	select {
	case <-measureResponseTime(url1):
		return url1, nil

	case <-measureResponseTime(url2):
		return url2, nil

	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func measureResponseTime(url string) chan struct{} {
	var ch = make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

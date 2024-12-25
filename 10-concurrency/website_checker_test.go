package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "any"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"google.com",
		"mayallo.com",
		"any",
	}

	want := map[string]bool{
		"google.com":  true,
		"mayallo.com": true,
		"any":         false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func mockSlowWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	websites := make([]string, 100)

	for i := 0; i < len(websites); i++ {
		websites[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(mockSlowWebsiteChecker, websites)
	}
}

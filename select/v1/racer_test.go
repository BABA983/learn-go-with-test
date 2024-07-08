package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.bing.com"
	fastURL := "http://127.0.0.1:5634"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

package racer

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	// a chan struct{} is the smallest data type available from a memory perspective
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

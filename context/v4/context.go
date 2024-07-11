package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	// Cancel()
}

func Server(s Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := s.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}

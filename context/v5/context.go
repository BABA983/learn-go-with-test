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
		data, err := s.Fetch(r.Context())
		if err != nil {
			return // todo: log error however you like
		}
		fmt.Fprint(w, data)
	}
}

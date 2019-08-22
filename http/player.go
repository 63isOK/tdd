package main

import (
	"fmt"
	"net/http"
)

// PlayerStore store the info(score etc)
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// StubPlayerStore is a store
type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

// PlayerServer is a server
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/palyers/"):]

	_, _ = fmt.Fprint(w, p.store.GetPlayerScore(player))
}

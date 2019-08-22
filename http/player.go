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

// GetPlayerScore return the score of the player
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.scores[name]
}

// PlayerServer is a server
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.processSocre(w, r)
	case http.MethodPost:
		p.processWin(w)
	}
}

func (p *PlayerServer) processSocre(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/palyers/"):]

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	_, _ = fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

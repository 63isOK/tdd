package main

import (
	"fmt"
	"net/http"
)

func getPlayerScore(name string) string {
	if name == "jim" {
		return "20"
	}
	if name == "tom" {
		return "10"
	}

	return ""
}

// PlayerServer start a server
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/palyers/"):]

	_, _ = fmt.Fprintf(w, getPlayerScore(player))
}

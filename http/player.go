package main

import (
	"fmt"
	"net/http"
)

// PlayerServer start a server
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/palyers/"):]

	if player == "jim" {
		_, _ = fmt.Fprintf(w, "20")
	}

	if player == "tom" {
		_, _ = fmt.Fprintf(w, "10")
	}
}

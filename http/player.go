package main

import (
	"fmt"
	"net/http"
)

// PlayerServer start a server
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "20")
}

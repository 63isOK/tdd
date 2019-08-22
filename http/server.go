package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatalf("listen on 8000 failed %v", err)
	}
}

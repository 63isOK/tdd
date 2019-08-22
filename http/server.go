package main

import (
	"log"
	"net/http"
)

type otherSotre struct{}

func (o *otherSotre) GetPlayerScore(name string) int {
	return 12
}

func main() {
	server := &PlayerServer{&otherSotre{}}
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		log.Fatalf("listen on 8000 failed %v", err)
	}
}

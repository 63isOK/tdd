package main

import (
	"fmt"
	"net/http"

	"github.com/63isOK/tdd/greet"
	"github.com/63isOK/tdd/integers"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	greet.Greet(w, "server")
}

func main() {
	fmt.Println(integers.Add(1, 3))

	http.ListenAndServe(":9000", http.HandlerFunc(greetHandler))
}

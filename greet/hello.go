package greet

import (
	"fmt"
	"io"
)

// Greet say hello to someone
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "hello %s", name)
}

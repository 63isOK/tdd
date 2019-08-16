package macking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord = "Go!"
	count     = 3
)

// Do do something
func Do(w io.Writer) {
	for i := count; i > 0; i-- {
		time.Sleep(1 * time.Second)
		_, _ = fmt.Fprintln(w, i)
	}
	time.Sleep(1 * time.Second)
	_, _ = fmt.Fprint(w, finalWord)
}

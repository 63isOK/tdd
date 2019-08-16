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

// Sleeper is interface
type Sleeper interface {
	Sleep()
}

// RealSleep is real sleep
type RealSleep struct {
	duration time.Duration
}

// Sleep is sleep
func (r *RealSleep) Sleep() {
	time.Sleep(r.duration)
}

// Do do something
func Do(w io.Writer, s Sleeper) {
	for i := count; i > 0; i-- {
		s.Sleep()
		_, _ = fmt.Fprintln(w, i)
	}
	s.Sleep()
	_, _ = fmt.Fprint(w, finalWord)
}

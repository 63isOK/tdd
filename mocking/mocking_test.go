package macking

import (
	"bytes"
	"testing"
)

func TestDo(t *testing.T) {
	buffer := &bytes.Buffer{}

	s := &spy{}
	Do(buffer, s)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("want '%s' got '%s'", want, got)
	}

	if s.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", s.Calls)
	}
}

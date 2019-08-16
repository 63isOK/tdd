package macking

import (
	"bytes"
	"testing"
)

func TestDo(t *testing.T) {
	buffer := &bytes.Buffer{}
	Do(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

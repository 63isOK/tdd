package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "tom")

	got := buffer.String()
	want := "hello tom"

	if got != want {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

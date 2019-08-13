package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld("Jim")
	want := "hello Jimm"
	if got != want {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "hello Tom"
	if got != want {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

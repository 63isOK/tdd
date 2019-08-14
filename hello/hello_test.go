package main

import "testing"

func TestHelloWorld(t *testing.T) {

	t.Run("say hello to someone", func(t *testing.T) {
		got := HelloWorld("Jim")
		want := "hello Jim"
		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	})

	t.Run("say hello world when there is nobody", func(t *testing.T) {
		got := HelloWorld("")
		want := "hello world"
		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	})
}

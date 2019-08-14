package main

import "testing"

func TestHelloWorld(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	}

	t.Run("say 你好 to someone", func(t *testing.T) {
		got := HelloWorld("Jim", chineseLanguage)
		want := "你好 Jim"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to someone", func(t *testing.T) {
		got := HelloWorld("Jim", "")
		want := "hello Jim"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when there is nobody", func(t *testing.T) {
		got := HelloWorld("", "")
		want := "hello world"
		assertCorrectMessage(t, got, want)
	})
}

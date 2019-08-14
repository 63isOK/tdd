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

	t.Run("say bonjour to someone", func(t *testing.T) {
		got := HelloWorld("Jim", frenchLanguage)
		want := "bonjour Jim"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello(english) to someone", func(t *testing.T) {
		got := HelloWorld("Jim", englishLanguage)
		want := "hello Jim"
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

func TestGetPrefix(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	}

	t.Run("get chinese prefix", func(t *testing.T) {
		got := getPrefix(chineseLanguage)
		want := chinesePrefix
		assertCorrectMessage(t, got, want)
	})

	t.Run("get french prefix", func(t *testing.T) {
		got := getPrefix(frenchLanguage)
		want := frenchPrefix
		assertCorrectMessage(t, got, want)
	})

	t.Run("get english prefix", func(t *testing.T) {
		got := getPrefix(englishLanguage)
		want := englishPrefix
		assertCorrectMessage(t, got, want)
	})

	t.Run("get prefix without specific language", func(t *testing.T) {
		got := getPrefix("")
		want := englishPrefix
		assertCorrectMessage(t, got, want)
	})
}

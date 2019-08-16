package main

import "testing"

func assertStrings(t *testing.T, want, got string) {
	t.Helper()

	if want != got {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "there is a test case"}

	t.Run("known word", func(t *testing.T) {
		got, _ := d.Search("test")
		want := "there is a test case"

		assertStrings(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := d.Search("test2")
		want := "there is no test case"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

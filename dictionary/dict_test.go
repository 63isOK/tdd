package main

import "testing"

func assertStrings(t *testing.T, want, got string) {
	t.Helper()

	if want != got {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

func TestSearch(t *testing.T) {
	d := map[string]string{"test": "there is a test case"}
	got := Search(d, "test")
	want := "there is a test case"

	assertStrings(t, want, got)
}

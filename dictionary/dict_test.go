package main

import "testing"

func TestSearch(t *testing.T) {
	d := map[string]string{"test": "there is a test case"}
	got := Search(d, "test")
	want := "there is a test case"

	if got != want {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

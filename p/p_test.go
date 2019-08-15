package main

import "testing"

func TestP(t *testing.T) {
	got := P(2.0, 2.0)
	want := 8.0

	if got != want {
		t.Errorf("want '%.2f' got '%.2f'", want, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(2, 3)
	want := 6.0

	if got != want {
		t.Errorf("want '%.2f' got '%.2f'", want, got)
	}
}

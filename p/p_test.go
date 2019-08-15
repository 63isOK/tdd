package main

import "testing"

func TestRectanglePerimeter(t *testing.T) {
	got := RectanglePerimeter(2.0, 2.0)
	want := 8.0

	if got != want {
		t.Errorf("want '%.2f' got '%.2f'", want, got)
	}
}

func TestRectangleArea(t *testing.T) {
	got := RectangleArea(2, 3)
	want := 6.0

	if got != want {
		t.Errorf("want '%.2f' got '%.2f'", want, got)
	}
}

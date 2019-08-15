package main

import "testing"

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{2, 3}
	got := RectanglePerimeter(r)
	want := 10.0

	if got != want {
		t.Errorf("want '%.2f' got '%.2f'", want, got)
	}
}

func TestRectangleArea(t *testing.T) {
	r := Rectangle{2, 3}
	got := RectangleArea(r)
	want := 6.0

	if got != want {
		t.Errorf("want '%.2f' got '%.2f'", want, got)
	}
}

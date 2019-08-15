package main

import (
	"fmt"
	"testing"
)

func float64Compare(x, y float64) bool {
	sx := fmt.Sprintf("%.2f", x)
	sy := fmt.Sprintf("%.2f", y)

	return sx == sy
}

func assertCorrectMessage(t *testing.T, want, got float64) {
	t.Helper()

	if !float64Compare(got, want) {
		t.Errorf("want '%.2f' got '%.2f'", want, got)
	}
}

func TestPerimeter(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{2, 3}
		got := RectanglePerimeter(r)
		want := 10.0

		assertCorrectMessage(t, want, got)
	})

	t.Run("circle", func(t *testing.T) {
		r := Circle{2}
		got := CirclePerimeter(r)
		want := 12.56

		assertCorrectMessage(t, want, got)
	})
}

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{2, 3}
		got := RectangleArea(r)
		want := 6.0

		assertCorrectMessage(t, want, got)
	})

	t.Run("circle", func(t *testing.T) {
		r := Circle{3}
		got := CircleArea(r)
		want := 28.26

		assertCorrectMessage(t, want, got)
	})
}

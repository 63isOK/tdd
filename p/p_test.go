package main

import (
	"testing"
)

func assertCorrectMessage(t *testing.T, want, got float64) {
	t.Helper()

	if int(want*100) != int(got*100) {
		t.Errorf("want '%.2f' got '%.2f'", want, got)
	}
}

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		assertCorrectMessage(t, want, got)
	}

	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{2, 3}
		want := 10.0

		checkPerimeter(t, &r, want)
	})

	t.Run("circle", func(t *testing.T) {
		r := Circle{2}
		want := 12.56

		checkPerimeter(t, &r, want)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectMessage(t, want, got)
	}

	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{2, 3}
		want := 6.0

		checkArea(t, &r, want)
	})

	t.Run("circle", func(t *testing.T) {
		r := Circle{3}
		want := 28.26

		checkArea(t, &r, want)
	})
}

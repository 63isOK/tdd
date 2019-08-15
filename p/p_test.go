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

	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{2, 3}, 10},
		{&Circle{2}, 12.56},
		{&Square{10}, 40},
		{&Square{3.32}, 13.28},
	}

	for _, tt := range areaTest {
		checkPerimeter(t, tt.shape, tt.want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectMessage(t, want, got)
	}

	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{3, 4}, 12},
		{&Circle{10}, 314},
		{&Square{10}, 100},
		{&Square{3.32}, 11.02},
	}

	for _, tt := range areaTest {
		checkArea(t, tt.shape, tt.want)
	}
}

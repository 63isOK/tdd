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
		{shape: &Rectangle{Width: 3, Height: 4}, want: 14},
		{shape: &Circle{Raduis: 10}, want: 62.8},
		{shape: &Square{Width: 10}, want: 40},
		{shape: &Square{Width: 3.32}, want: 13.28},
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
		{shape: &Rectangle{Width: 3, Height: 4}, want: 12},
		{shape: &Circle{Raduis: 10}, want: 314},
		{shape: &Square{Width: 10}, want: 100},
		{shape: &Square{Width: 3.32}, want: 11.02},
	}

	for _, tt := range areaTest {
		checkArea(t, tt.shape, tt.want)
	}
}

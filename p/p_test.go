package main

import (
	"testing"
)

func assertCorrectMessage(t *testing.T, want, got float64, shape Shape) {
	t.Helper()

	if int(want*100) != int(got*100) {
		t.Errorf("%#v want '%.2f' got '%.2f'", shape, want, got)
	}
}

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		assertCorrectMessage(t, want, got, shape)
	}

	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: &Rectangle{Width: 3, Height: 4}, want: 14},
		{name: "circle", shape: &Circle{Raduis: 10}, want: 62.8},
		{name: "square", shape: &Square{Width: 10}, want: 40},
		{name: "square", shape: &Square{Width: 3.32}, want: 13.28},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.shape, tt.want)
		})
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectMessage(t, want, got, shape)
	}

	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: &Rectangle{Width: 3, Height: 4}, want: 12},
		{name: "circle", shape: &Circle{Raduis: 10}, want: 314},
		{name: "square", shape: &Square{Width: 10}, want: 100},
		{name: "square", shape: &Square{Width: 3.32}, want: 11.02},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}

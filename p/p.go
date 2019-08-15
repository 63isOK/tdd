package main

import (
	"fmt"
)

// Rectangle is a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// PI is pi
const PI = 3.14

// RectanglePerimeter calc the perimeter of the rectangle
func RectanglePerimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// RectangleArea calc the area fo the rectangle
func RectangleArea(r Rectangle) float64 {
	return r.Width * r.Height
}

// Circle is a circle
type Circle struct {
	Raduis float64
}

// CirclePerimeter calc the perimeter of the circle
func CirclePerimeter(c Circle) float64 {
	return 2 * PI * c.Raduis
}

// CircleArea calc the area fo the circle
func CircleArea(c Circle) float64 {
	return PI * c.Raduis * c.Raduis
}

func main() {
	c := Circle{2}
	fmt.Println(CircleArea(c))
}

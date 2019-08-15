package main

import (
	"fmt"
)

// Rectangle is a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// RectanglePerimeter calc the perimeter of the rectangle
func RectanglePerimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// RectangleArea calc the area fo the rectangle
func RectangleArea(r Rectangle) float64 {
	return r.Width * r.Height

}

func main() {
	r := Rectangle{1, 2}
	fmt.Println(RectanglePerimeter(r))
}

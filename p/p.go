package main

import (
	"fmt"
)

// RectanglePerimeter calc the perimeter of the rectangle
func RectanglePerimeter(w, h float64) float64 {
	return 2 * (w + h)
}

// RectangleArea calc the area fo the rectangle
func RectangleArea(w, h float64) float64 {
	return w * h

}

func main() {
	fmt.Println(RectanglePerimeter(1, 2))
}

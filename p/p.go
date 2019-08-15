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

// Perimeter calc the perimeter of the rectangle
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area calc the area fo the rectangle
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle is a circle
type Circle struct {
	Raduis float64
}

// Perimeter calc the perimeter of the circle
func (c *Circle) Perimeter() float64 {
	return 2 * PI * c.Raduis
}

// Area calc the area fo the circle
func (c *Circle) Area() float64 {
	return PI * c.Raduis * c.Raduis
}

func main() {
	c := Circle{2}
	fmt.Println(c.Area())
}

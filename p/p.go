package main

import (
	"fmt"
)

// Shape abstract shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

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

// Square is a square
type Square struct {
	Width float64
}

// Perimeter calc the perimeter of the square
func (s *Square) Perimeter() float64 {
	return 4 * s.Width
}

// Area calc the area fo the square
func (s *Square) Area() float64 {
	return s.Width * s.Width
}

func main() {
	c := Circle{2}
	fmt.Println(c.Area())
}

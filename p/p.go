package main

import (
	"fmt"
)

// P calc the perimeter of the rectangle
func P(w, h float64) float64 {
	return 2 * (w + h)
}

// Area calc the area fo the rectangle
func Area(w, h float64) float64 {
	return w * h

}

func main() {
	fmt.Println(P(1, 2))
}

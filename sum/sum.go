package main

import (
	"fmt"
)

// Sum calc the sum of array
func Sum(n [5]int) (sum int) {
	for _, x := range n {
		sum += x
	}

	return
}

// SumSlice calc the sum of slice
func SumSlice(n []int) (sum int) {
	for _, x := range n {
		sum += x
	}

	return
}

func main() {
	s := Sum([5]int{1, 1, 1, 1, 1})
	fmt.Println(s)
}

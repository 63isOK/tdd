package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("want '%d' got '%d', %v", want, got, numbers)
	}
}

func TestSumSlice(t *testing.T) {

	type params struct {
		input []int
		want  int
	}

	cases := []params{
		{[]int{1, 2, 3}, 6},
		{[]int{}, 0},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, 2, 3, 1, 2, 3, 2, 3}, 18},
		{[]int{1, 2, 3, 4, 5}, 15},
	}

	for _, x := range cases {
		t.Run("test sum of slice", func(t *testing.T) {
			got := SumSlice(x.input)
			if got != x.want {
				t.Errorf("want '%d' got '%d', %v", x.want, got, x.input)
			}
		})
	}
}

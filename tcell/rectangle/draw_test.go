package main

import (
	"testing"
)

func TestGenData(t *testing.T) {
	a := genData(10, 12)
	if a[3][2] != 3 ||
		a[9][9] != 9 ||
		a[11][9] != 11 ||
		a[1][3] != 3 {
		t.Fatal("data is error")
	}
}

package test

import "testing"

func add(x int, y int) int {
	if x >= 0 && y >= 0 {
		return x + y
	}
	return 0

}

func Test_add(t *testing.T) {
	sum := add(1, 3)
	if sum != 5 {
		t.Error("wrong sum")
	}

}

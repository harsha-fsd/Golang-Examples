package test

import "testing"

func add(x int, y int) int {
	return x + y
}

func Test_add(t *testing.T) {
	sum := add(1, 3)
	if sum != 4 {
		t.Error("wrong sum")
	}

}

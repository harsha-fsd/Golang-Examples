package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	y := 0
  // comment out 2 of the following line to test any func
	fmt.Println(safeDiv(x, y))
	//program panics (Exception is thrown) after this call as divide by zero error occurs
	fmt.Println(div(x, y))
    check(4)
}

func div(x int, y int) int {
	return x / y
}

//try catch style golang way
func safeDiv(x int, y int) int {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	return x / y
}


func check(x int) {
	if x < 5 {
		panic("Illegal Argument: " + strconv.Itoa(x))
	}
	fmt.Printf("value %d\n is valid", x)
}

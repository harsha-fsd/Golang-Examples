package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println("Value of x:", x)
	var xPointer *int = &x
	fmt.Println("X is at location in memory:", xPointer)
	fmt.Println("Value of x on deferencing:", *xPointer)
	changeIt(x)
	fmt.Println("Value of x after calling changeIt func:", x)
	changeItByPtr(xPointer)
	fmt.Println("Value of x after calling changeItByPtr func:", x)

}

func changeIt(x int) {
	x = x * 2
}

func changeItByPtr(x *int) {
	*x = (*x) * 2

}

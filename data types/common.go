package main

import "fmt"

func main() {
	a := 5
	b := -5
	c := 99.99
	d := "Hello Go World!"
	e := false
	f := []int{1, 2, 3, 4}
	g := map[string]float64{"strawberry": 3.2}
	const ZERO = 0

	// %T prints type of the variable
	//alternatively Type
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", c)
	fmt.Printf("d: %T\n", d)
	fmt.Printf("e: %T\n", e)
	fmt.Printf("f: %T\n", f)
	fmt.Printf("g: %T\n", g)
	fmt.Println(ZERO)

}

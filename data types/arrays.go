package main

import (
	"fmt"
	"sort"
)

func main() {
	//array
	a := [2]int{}
	a[0] = 1
	a[1] = 2
	fmt.Println(a)
	b := [...]int{5, 6, 7}
	fmt.Println("Size of b:", len(b))
	//slice
	slice := []int{}
	slice = append(slice, 19)
	fmt.Println(slice[0])
	//creates a slice with 2 elements with default init values
	slice = make([]int, 2)

	slice[1] = -190
	fmt.Println("slice:", slice)
	//sub slice
	sub := slice[1:]
	fmt.Println("subslice:", sub)
	//add an element to slice
	slice = append(slice, 350)
	for i, v := range slice {
		fmt.Printf("slice value at index %d is %d", i, v)
		fmt.Println()
	}

	fmt.Println("slice original:", slice)
	sort.Ints(slice)
	fmt.Println("slice sorted:", slice)

}

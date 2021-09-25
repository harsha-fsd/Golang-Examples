package main

import (
	"fmt"
	"sort"
)

type Grade struct {
	Name string
	Rank int
}

func main() {
	x := 5
	double, square := doubleAndSquare(x)
	fmt.Printf("Double of %d: %d\n", x, double)
	fmt.Printf("Square of %d: %d\n", x, square)
	grades := []Grade{
		{
			Name: "Cecilia",
			Rank: 2,
		},
		{
			Name: "Alice",
			Rank: 3,
		},
		{
			Name: "Barbara",
			Rank: 1,
		},
	}
	fmt.Println("Unsorted Grades: ", grades)

	//passing function to a function
	sort.Slice(grades, func(i, j int) bool {
		return grades[i].Rank < grades[j].Rank
	})
	fmt.Println("Sorted Grades: ", grades)

	retFunc1 := retFunc(-1)
	fmt.Println(retFunc1(5))
	retFunc2 := retFunc(1)
	fmt.Println(retFunc2(5))
	fmt.Println(add(10, 2, -4, 5, 6))
	deferredFunc()
}

//returns multiple values
func doubleAndSquare(x int) (double int, square int) {
	return x * 2, x * x
}

//func returns another func based on input
func retFunc(x int) func(int) int {
	if x < 0 {
		return func(y int) int {
			return 2 * y
		}
	}
	return func(y int) int {
		return y * y
	}

}

//variable length arguments example
func add(numbers ...int) int {
	sum := 0
	if len(numbers) > 0 {

		for i := range numbers {
			sum += numbers[i]
		}
	}
	return sum
}

//defer keyword function
func deferredFunc() {
	for i := 1; i <= 5; i++ {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	//prints this line first followed by 5 4 3 2 1 (LIFO)
	fmt.Println("Func execution deferred")
}

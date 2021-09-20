package main

import "fmt"

// global variables can be declared/initialized by only using var keyword not short hand assignment
var global = 20

func main() {
	//variable declaration using var keyworkd. initialized to zero value by default
	var i int
	//short hand assignment of the variable
	j := 5
	fmt.Println(i + j)
	fmt.Println(global + i + j)

}

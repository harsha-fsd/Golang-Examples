package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	//while loop style
	i := 1
	for {
		fmt.Print(i, " ")
		i++
		if i>10{
			break
		}
	}
	fmt.Println()

}

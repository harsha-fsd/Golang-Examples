package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argument := os.Args[1]
	x, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Printf("Can't '%s' convert to int", argument)
		fmt.Println()
		return
	}
	if x <= 5 {
		fmt.Println("value is less than or equal to 5")
	} else if x >= 100 {
		fmt.Println("value is less than or equal to 100")
	} else {
		fmt.Println("value is ", x)
	}
	switch argument {
	case "0":
		fmt.Println("Zero")
	case "1", "2", "3", "4", "5":
		fmt.Println("Value is between 1 and 5")
		fallthrough
	default:
		fmt.Println("Value: ", argument)
	}
}

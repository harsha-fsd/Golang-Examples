package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	err := check(5, 6)
	print(err)
	err = check(0, 0)
	print(err)
	//srtconv.Atoi converts from string to int if valid otherwise returns error
	fmt.Println(strconv.Atoi("12d"))
}

func check(a int, b int) error {
	if a == 0 && b == 0 {
		return errors.New("Both the arguments can't be zero")
	}
	return nil
}

func print(err error) {
	if err == nil {
		fmt.Println("All checks passed. No errors")
	} else {
		fmt.Println("Error: ", err.Error())
	}
}

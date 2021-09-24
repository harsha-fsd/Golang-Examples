package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Current time: ", start)
	date, err := time.Parse("02 January 2006", "24 September 2021")
	if err == nil {
		fmt.Println("Formatted date is ", date)
	}
	epoch := time.Now().Unix()
	fmt.Println("Epoch time is ", epoch)
	fmt.Printf("Type of epoch is %T\n", epoch)
	duration := time.Since(start)
	fmt.Println("Time since start of this program in microseconds:", duration)
	loc, _ := time.LoadLocation("GMT")
	fmt.Println("Current GMT time is:", time.Now().In(loc))
}

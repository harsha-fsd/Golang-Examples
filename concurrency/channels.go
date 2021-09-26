package main

import (
	"fmt"
	"time"
)

func main() {
	// A go channel allows go routines to exchange data
	c := make(chan int, 10)
	writeToChannel(c, 10)
	// reads the first element of channel
	fmt.Println(readFromChannel(c), " ")

	// reads the next nine elements of the channel and adds to 1 to each of the elements
	for i := range c {
		fmt.Print(i+1, " ")
	}
	fmt.Println()

	c1 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("Timeout c1")
	}
}

func writeToChannel(c chan int, x int) {
	for i := 1; i <= x-1; i++ {
		c <- i
	}
	close(c)

}

func readFromChannel(c chan int) int {
	return <-c
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//functions with 'go' keyword indicates go routines
	go func(x int) {
		fmt.Println("Executing another go routine in main func...")
		fmt.Printf("%d\n", x)
		fmt.Println("Exiting another go routine in main func...")
	}(1258)
	fmt.Println("Executing main go routine....")
	//wait for 100ms for go routine to finish
	time.Sleep(time.Duration(100000))

	//better waiting with waitgroups
	waitGroupsTest()
	fmt.Println("Exiting main routine.....")
}

func waitGroupsTest() {
	waitGrp := sync.WaitGroup{}
	waitGrp.Add(1)
	go func(x int) {
		defer waitGrp.Done()
		fmt.Println("Executing a go routine inside waitGroupsTest func....")
		fmt.Printf("%d\n", x)
		fmt.Println("Exiting a go routine inside waitGroupsTest func....")
	}(1258)
	waitGrp.Wait()
}

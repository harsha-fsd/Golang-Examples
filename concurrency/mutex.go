package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex

//shared variable among go routines
var x int

func main() {
	var wg sync.WaitGroup
	fmt.Printf("%d ", read())
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			change(i)
			fmt.Printf("->%d", read())
		}(i)
	}
	wg.Wait()
	fmt.Printf("->%d\n", read())
}

func change(i int) {
	m.Lock()
	x = i
	m.Unlock()
}

func read() int {
	m.Lock()
	v := x
	m.Unlock()
	return v
}

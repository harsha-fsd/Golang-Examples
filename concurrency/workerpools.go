package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Client struct {
	id    int
	input int
}

type Result struct {
	job    Client
	square int
}

var size = runtime.GOMAXPROCS(runtime.NumCPU())
var clients = make(chan Client, size)
var data = make(chan Result, size)

func main() {
	go create(100)
	//for blocking
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d", d.job.id)
			fmt.Println()
			fmt.Printf("Square of %d (Result): %d", d.job.input, d.square)
			fmt.Println()
		}
		finished <- true
	}()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)
	fmt.Printf("Finished: %v\n", <-finished)
}

func worker(wg *sync.WaitGroup) {
	for c := range clients {
		fmt.Println("Processing client Id", c.id)
		square := c.input * c.input
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

func create(n int) {
	for i := 0; i <= n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

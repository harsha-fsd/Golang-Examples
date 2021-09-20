package main

import "fmt"

func main() {
	fmt.Printf("Please enter your name: ")
	var name string
	//Read user input
	fmt.Scanln(&name)
	fmt.Println("Your name is ", name)

}

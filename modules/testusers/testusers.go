package main

import (
	"fmt"
	"users"
)

func main() {
	u := users.User{
		Name: "Test",
		ID:   "9999",
	}
	fmt.Println("User Details: ", u.GetUserDetails())
}

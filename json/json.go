package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	//adding json tags to the fields
	Name string `json:"name"`
	ID   int    `json:"userId"`
}

func main() {
	user := User{Name: "Jeff Bezoz", ID: 999}
	userJson, err := json.Marshal(&user)
	if err == nil {
		fmt.Printf("User Json %s\n", userJson)
	} else {
		fmt.Println("Error marshalling user", err)
	}
	unMarshalUser := User{}
	err = json.Unmarshal(userJson, &unMarshalUser)
	if err == nil {
		fmt.Println("User struct from userJson", unMarshalUser)
	} else {
		fmt.Println("Error while trying to unmarshal userJson", err)
	}
}

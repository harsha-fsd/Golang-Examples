package main

import (
	"fmt"
)

type Entry struct {
	FirstName    string
	LastName     string
	MobileNumber string
}

//slice same as
var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.LastName == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	data = append(data, Entry{"James", "Miller", "9988664455"})
	data = append(data, Entry{"Maria", "Garcia", "8899664455"})
	data = append(data, Entry{"Sarah", "Jones", "7799664455"})
	list()
	entry := search("Jones")
	fmt.Println(entry.FirstName)

}

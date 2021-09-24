package main

import "fmt"

func main() {
	m := map[string]int{}
	m["c"] = 3
	m["a"] = 1
	m["b"] = 2
	m["e"] = 5
	m["d"] = 4
	fmt.Println(m["c"])
	for k, v := range m {
		fmt.Println("Key, Value:", k, " ", v)

	}
	//get all keys of the map
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println("Keys of the map:", keys)
	//clears all contents of the map
	for k := range m {
		delete(m, k)
	}
	fmt.Println(m)

}

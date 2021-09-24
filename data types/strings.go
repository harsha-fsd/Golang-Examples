package main

import (
	"fmt"
	s "strings"
	"unicode"
)

func main() {
	str := "banana"
	fmt.Printf("str is a %T \n", str)
	str = "work hard"
	for _, c := range str {
		fmt.Printf("%c ", c)

	}
	fmt.Println()
	fmt.Println(s.HasSuffix("apple", "ap"))
	fmt.Println(s.Index("apple", "ap"))
	fmt.Println(s.ToUpper("apple"))
	fmt.Println(s.Contains("apple", "banana"))
	fmt.Println(s.EqualFold("test", "Test"))
	fmt.Println("test" == "Test")

	trimFunc := func(c rune) bool {
		return unicode.IsDigit(c)
	}
	fmt.Println(s.TrimFunc("1Apple123", trimFunc))

}

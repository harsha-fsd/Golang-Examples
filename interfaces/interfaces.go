package main

import (
	"fmt"
	"reflect"
)

type Shape interface {
	draw()
}

type Circle struct {
}

type Triangle struct {
}

func (c Circle) draw() {
	fmt.Println("Circle.....")

}

func (t Triangle) draw() {
	fmt.Println("Triangle.....")
}

func main() {
	c := Circle{}
	t := Triangle{}
	c.draw()
	t.draw()
	_, ok := interface{}(c).(Shape)
	if ok {
		fmt.Println(reflect.TypeOf(c), "is a Shape")
	}
}

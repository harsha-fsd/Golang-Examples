package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	FirstName string
	LastName  string
	ID        int
}

//type methods

func (p *Person) SetFirstName(s string) {
	p.FirstName = s
}

func (p *Person) SetLastName(s string) {
	p.LastName = s

}
func (p Person) Name() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	p := new(Person)
	p.SetFirstName("Donald")
	p.SetLastName("Trump")

	fmt.Println("p: ", p)
	fmt.Println("p name: ", p.Name())
	p1 := Person{
		FirstName: "Donald",
		LastName:  "Trump",
	}
	//one level comparsion
	fmt.Println("p equals to p1?", *p == p1)
	//use reflection to check equality
	fmt.Println("p equals to p1 using reflection?", reflect.DeepEqual(*p, p1))
	p.ID = 1
	p1.ID = 2
	fmt.Println("p equals to p1 with ids?", *p == p1)
	// slice of persons
	persons := []Person{}
	persons = append(persons, p1)
	persons = append(persons, *p)
	fmt.Println(persons)
	p2 := Person{}
	//using reflection to set values of p2
	r := reflect.ValueOf(&p2).Elem()
	fmt.Println("Type of p2", r.Type())
	typeOfP2 := r.Type()
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		fmt.Println("Field ", field.Kind())
		if typeOfP2.Field(i).Name == "LastName" {
			field.SetString("Trump")
		} else if typeOfP2.Field(i).Name == "FirstName" {
			field.SetString("Donald")
		} else if typeOfP2.Field(i).Name == "ID" {
			field.SetInt(3)
		}
	}
	fmt.Println("p2 after setting values by reflection", p2)

}

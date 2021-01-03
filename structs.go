package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 10
	return &p
}

func main() {
	fmt.Println(newPerson("Jon"))
	fmt.Println(person{age: 29, name: "My Name"})
}

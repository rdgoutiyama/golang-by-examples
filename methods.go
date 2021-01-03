package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) setWidthWithoutReference() {
	r.width = 10
}

func (r *rect) setWidthWithReference() {
	r.width = 1
}

func main() {
	rect := rect{width: 10, height: 5}
	fmt.Println(rect.area())

	rect.setWidthWithoutReference()
	fmt.Println(rect.area())

	rect.setWidthWithReference()
	fmt.Println(rect.area())
}

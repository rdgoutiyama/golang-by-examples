package main

import "fmt"

func vals() (int, int) {
	return 10, 11
}

func main() {
	a, b := vals()

	fmt.Println(a, b)
}

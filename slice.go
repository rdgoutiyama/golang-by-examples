package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	fmt.Println(s)

	s = append(s, "c")
	fmt.Println(s)

	t := make([]string, 1)
	copy(t, s)
	fmt.Println(t)

	l := s[:2]
	fmt.Println(l)
}

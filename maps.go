package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 2
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)

	delete(m, "k2")
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println(prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

	fmt.Println(n["foo"])
}

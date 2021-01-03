package main

import "fmt"

func zeroVal(val int) {
	val = 0
}

func zeroPtrVal(val *int) {
	*val = 0
}

func main() {
	i := 1

	zeroVal(i)
	fmt.Println(i)

	zeroVal(i)
	fmt.Println(i)

	zeroPtrVal(&i)
	fmt.Println(i)

	zeroVal(i)
	fmt.Println(i)

}

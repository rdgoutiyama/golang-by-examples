package main

import "fmt"

func sum(nums ...int) int {

	total := 0

	for _, num := range nums {
		fmt.Println(num)
		total += num
	}

	return total

}

func main() {
	fmt.Println(sum(1, 2))
}

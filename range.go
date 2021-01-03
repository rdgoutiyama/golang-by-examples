package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(nums)
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index of ", num, " is ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Println(k, " - ", v)
	}
}

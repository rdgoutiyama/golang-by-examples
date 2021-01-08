package main

import (
	"fmt"
	"strings"
)

func worker(values []string, c chan string) {
	c <- "values: " + strings.Join(values, "")
}

func main() {

	c := make(chan string)
	nums := []string{"1", "2", "A", "B", "C"}
	go worker(nums[:len(nums)/2], c)
	go worker(nums[len(nums)/2:], c)

	fmt.Println(<-c)
	fmt.Println(<-c)

	// time.Sleep(time.Second)

}

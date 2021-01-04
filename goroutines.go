package main

import (
	"fmt"
	"time"
)

func f(description string) {
	for i := 0; i < 10; i++ {
		fmt.Println(description, i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// if comment next line, goroutine's functions will not be executed, because async call
	time.Sleep(time.Second)

	fmt.Println("Done")
}

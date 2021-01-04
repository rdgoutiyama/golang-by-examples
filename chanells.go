package main

import (
	"fmt"
	"time"
)

func greet(c chan string) {
	name := <-c
	fmt.Println("Hello ", name)
}

func main() {

	// messages := make(chan string)

	// go func() { messages <- "ping" }()

	// msg := <-messages
	// fmt.Println(msg)

	c := make(chan string)

	go greet(c)

	c <- "World"

	close(c)

	time.Sleep(time.Second)

}

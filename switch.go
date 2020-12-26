package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	description := ""

	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		description = "One"
	case 2:
		description = "Two"
	case 3:
		description = "Three"
	}

	fmt.Println(description)

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's AM")
	default:
		fmt.Println("It's PM")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Dont't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI("true")
}

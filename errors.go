package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

// argError implementa a interface de Errors
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg1 int) (int, error) {
	if arg1 == 42 {
		return -1, errors.New("número 42 não é possível ser usado")
	}

	return arg1 + 3, nil
}

func f2(arg int) (int, error) {

	if arg == 42 {
		return -1, &argError{arg, "número 42 não é possível ser usado"}
	}

	return arg, nil
}

func main() {

	for _, i := range []int{7, 12, 42} {
		fmt.Println(i)
		_, errorValue := f1(i)
		if errorValue != nil {
			fmt.Println(errorValue)
		}
	}

	fmt.Println("------------------------------")

	for _, i := range []int{7, 12, 42} {
		fmt.Println(i)
		_, errorValue := f2(i)
		if errorValue != nil {
			fmt.Println(errorValue)
		}
	}

	fmt.Println("------------------------------")
	returnedValue, e := f2(42)
	fmt.Println(returnedValue)
	fmt.Println(e.(*argError))

	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}

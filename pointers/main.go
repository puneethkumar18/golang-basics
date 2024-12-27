package main

import "fmt"

func add(a *int, b int) int {
	return *a + b
}

func main() {
	fmt.Println("We are starting Pointers in Go")

	var num int = 100
	var numPointer *int = &num

	fmt.Println("Value of num is ", num)
	fmt.Println("Address of num is ", numPointer)

	var pointer *string

	if pointer == nil {
		fmt.Println("Pointer is not nil")
	} else {
		fmt.Println("Pointer is nil")
	}

	fmt.Println("Value of pointer is ", *numPointer)

	fmt.Println("Addition of 100 and 20 is ", add(numPointer, 20))

}

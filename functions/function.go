package main

import "fmt"

func simplefunction() {
	fmt.Println("Simple function")
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

func multiply(a, b int) int {
	return a * b
}

func subtraction(a, b int) (result int) {
	result = a - b
	return
}

func division(a, b int) (result int) {
	result = a / b
	return
}

func main() {
	fmt.Println("We Are Learning Function In Go Lang")

	sum := add(3, 2)
	fmt.Println("Sum of 3,2 is ->", sum)

	mul := multiply(3, 2)
	fmt.Println("Multiply of 3,2 is -> ", mul)

	sub := subtraction(3, 2)
	fmt.Println("Subtraction of 3 and 2 is -> ", sub)

	divi := division(4, 2)
	fmt.Println("Division of 4 by 2 is ->", divi)

}

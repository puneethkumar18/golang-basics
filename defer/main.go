package main

import "fmt"

func initfunc() {
	fmt.Println("We are starting Defer in Go")
}
func end() {
	fmt.Println("We are ending Defer in Go")
}

func add(a, b int) int {
	return a + b
}

func main() {
	initfunc()
	defer end()

	fmt.Println("My Name is Puneeth Kumar")

	var a, b int
	fmt.Println("Enter the value of a and b")
	fmt.Scan(&a, &b)
	defer fmt.Println("Sum of a and b is:", add(a, b))

	fmt.Println("After printing sum")

	fmt.Println("Defer is used to delay the execution of a statement until the end of the function")
	fmt.Println("Defer Follows Last In First Out Rule")

}

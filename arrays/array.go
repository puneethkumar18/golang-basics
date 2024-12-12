package main

import "fmt"

func main() {
	fmt.Println("We are learning the Arrays Concepts")

	// Declaration
	var numbers = [10]int{1, 2, 2, 3, 4, 5, 6, 10, 2, 1}

	// this is not possilble
	//var ages[5]int = {21,22,23,89,90}
	var names [5]string
	names[0] = "Puneeth"
	names[1] = "Kumar"
	names[2] = "G"

	fmt.Println("All the names of the array is -> ", names)
	fmt.Println("All the Numbers in number arrray is ->", numbers)

	// to Print The Length of the array is
	fmt.Println("Length of the numbers Array is -> ", len(numbers))
	fmt.Println("Length of the names Array is -> ", len(names))

	// Empty array
	var lists [5]string

	fmt.Println("Lists array contents is ->", lists)
	lists[2] = "Puneeth"
	fmt.Printf("Printing the lists array with double quotes %q\n", lists)
}

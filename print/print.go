package main

import "fmt"

func main() {
	var name string = "Puneeth"
	var age int = 21
	var height float32 = 5.10

	//print usning Println which make every line break
	fmt.Println("My name is ->", name)
	fmt.Println("My age is ->", age)
	fmt.Println("My height is ->", height)
	fmt.Println("Name : ", name, "Age : ", age, "height : ", height)

	// using Printf
	fmt.Printf("Name is -> %s \t", name)
	fmt.Printf("Age is -> %d\t", age)
	fmt.Printf("Height is -> %.2f\n", height)

	// print Type Of the variable
	fmt.Printf("Name type is -> %T\n", name)
	fmt.Printf("Age type is -> %T\n", age)
	fmt.Printf("Height type is -> %T\n", height)
}

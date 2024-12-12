package main

import "fmt"

func main() {
	var name string = "Puneeth"
	fmt.Println("My name is -> ", name)

	var age int = 21
	fmt.Println("My age is -> ", age)

	var amIGraduated bool = true
	fmt.Println("Am I Graduated -> ", amIGraduated)

	var cgpa float32 = 7.02
	fmt.Println("My Graduation CGPA is -> ", cgpa)

	var description = "I am Persuing Graduation in JSS Science and technology Mysore"
	fmt.Println("My Description is -> ", description)

	const height = 5.10
	//  height = 6.1  -> constants can't be reassigned

	var weight = 70
	fmt.Println("My weight before changes -> ", weight)
	weight = 60
	fmt.Println("My weight After changes -> ", weight)

	// PublicVariable -> which can used outside the file or program(By starting variable name with CAP letter)
	var Public = "This is the variable Can be used outside the File"
	fmt.Println("Public Variable -> ", Public)

	// privateVariable -> Which Start with SMALL letter And Can't be use outside the File
	var private = "This is the variable Can't be usec outside the File"
	fmt.Println("private Variable -> ", private)

	// we Can decalere variavle with out var const tag
	Person := "Puneeth kumar G"
	fmt.Println("Another Way of decalaring -> ", Person)
}

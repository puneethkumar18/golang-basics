package main

import "fmt"

func main() {
	fmt.Println("We Are Starting Maps in Go")

	// Map declaration

	var map1 = map[string]int{"Puneeth": 1, "Raj": 2, "Ravi": 3}

	var map2 = map[string]int{"Puneeth": 1, "Raj": 2, "Ravi": 3}
	// using make function
	var myMap = make(map[string]int)
	myMap["Kumar"] = 1

	fmt.Println(myMap, map1, map2)

}

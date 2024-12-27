package main

import "fmt"

func main() {
	fmt.Println("We Are Starting Maps in Go")

	// Map declaration

	var map1 = map[string]int{"Puneeth": 1, "Raj": 2, "Ravi": 3}

	map1["Abhi"] = 4

	var map2 = map[string]int{"Puneeth": 10, "Raj": 2, "Ravi": 3, "Abhi": 10000}

	map2["Ramesh"] = 420
	// using make function
	var myMap = make(map[string]int)
	myMap["Kumar"] = 1

	fmt.Println(myMap, map1, map2)

	for key, value := range map2 {
		fmt.Println("key is ", key, "Value is ", value)
	}

	for key, value := range map1 {
		fmt.Println("key is ", key, "Value is ", value)
	}

}

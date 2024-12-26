package main

import "fmt"

func main() {
	fmt.Println("Eveything is perfectctly all right")

	//Slices declaration
	var name = []string{}
	fmt.Println("name", name)
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	numbers = append(numbers, 10, 12, 122)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d -> %d\n", i, numbers[i])
	}

	size := len(numbers)
	fmt.Println("Length", size)

	number := make([]int, 5, 10)
	number = append(number, 10, 20, 30, 40, 50, 60)

	for i := 0; i < len(number); i++ {
		fmt.Println("Index -> ", i, "Value -> ", number[i])
	}
	fmt.Println("first ele : ", number[0])
	fmt.Println("Size : ", len(number))
	fmt.Println("Capacity : ", cap(number))

}

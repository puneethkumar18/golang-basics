package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("We are starting Data Conversion in Go")

	var num int = 43

	fmt.Printf("Value of num is %v\ntype is %T\n", num, num)

	var data float64 = float64(num)

	fmt.Printf("Value of num is %v\ntype is %T\n", data, data)

	number_string := "1234"

	fmt.Printf("Value of number as String is %v\ntype is %T\n", number_string, number_string)

	number_int, _ := strconv.Atoi(number_string)
	fmt.Printf("Value of number as Int is %v\ntype is %T\n", number_int, number_int)
}

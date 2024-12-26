package main

import "fmt"

func main() {

	fmt.Println("We Are Starting Switch Statemets")

	var num int
	fmt.Println("Enter the number with in 7 i will give you the day")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("Wednesday")

	case 4:
		fmt.Println("Thursday")

	case 5:
		fmt.Println("Friday")

	case 6:
		fmt.Println("Saturday")

	default:
		fmt.Println("Sunday")

	}

	var temprature int

	fmt.Println("Enter The Tempature")

	fmt.Scan(&temprature)

	switch {
	case temprature < 30:
		fmt.Println("Its Hot")
	case temprature < 100:
		fmt.Println("Its Very Hot")

	default:
		fmt.Println("I don't know")
	}
}

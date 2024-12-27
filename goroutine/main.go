package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World!1")

	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Hello World!2")
}

func print() {
	fmt.Println("Printing the required Function")
}
func main() {
	fmt.Println("Learning GORoutine....")

	go sayHello()

	print()

	time.Sleep(time.Millisecond * 3000)
}

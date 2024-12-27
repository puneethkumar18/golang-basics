package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("We are starting Time in Go")

	// time package is used to work with time in Go
	currentTime := time.Now()

	fmt.Println("Current Time is:", currentTime)

	fmt.Printf("Type of the Current Time is %T\n", currentTime)

	formatted := currentTime.Format("2006-01-02 Monday 3:01 PM")
	fmt.Println("Formatted Time is:", formatted)

}

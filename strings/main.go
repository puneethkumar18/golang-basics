package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("We are starting Strings in Go")

	data := "banana,apple,orange,grapes"
	parts := strings.Split(data, ",")

	fmt.Printf("Value of data is %v\ntype is %T\n", data, data)
	fmt.Printf("Value of parts is %v\ntype is %T\n", parts, parts)

	var numbers string = "one,two,three,four,one,two,three,four,one"

	fmt.Println("Number of ones in numbers string are :", strings.Count(numbers, ","))

	fmt.Println("Parts of data are:", strings.Index(data, "apple"))

	str := "      pueeth kumar       "

	fmt.Println("Befor Trim:", str)
	fmt.Println("before Uppercse :", str)
	str = strings.Trim(str, " ")
	fmt.Println("After Trim:", str)
	fmt.Println("After Uppercse :", strings.ToUpper(str))

}

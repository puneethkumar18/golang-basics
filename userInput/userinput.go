package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Scan Using {Scanf}
	fmt.Println("This Is The USER INPUT FILE")
	//var name string
	// fmt.Scanf("%s", &name)
	// fmt.Println("Hello ,", name)

	// Scan using Scan
	// fmt.Println("Tell Me Your Name -> ")
	// fmt.Scan(&name)
	// fmt.Println("Your Name is ->", name)

	// using bufio package
	var setence string
	reader := bufio.NewReader(os.Stdin)
	setence, _ = reader.ReadString('\n')
	fmt.Println("Your have printed this -> ", setence)
}

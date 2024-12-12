package main

import (
	"fmt"
	"go_lang/myutil"
)

func main() {
	fmt.Println("Learn Go Lang BY Myself")
	sum()
	myutil.PrintNumber()
	fmt.Println(myutil.Global)

}

func sum() {
	fmt.Println("This is my Sumn program")
}

package main

import "fmt"

func division(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divisor is Zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("We have started error handling!")

	res, _ := division(10, 2)
	// here underscore is used to get that error but we dont want that to use in  any where
	fmt.Println("Division of 10 by 2 is -> ", res)

	ans, err := division(90, 0)
	if err != nil {
		fmt.Println("Error Occured is -> ", err)
	} else {
		fmt.Println("Division of 90 / 0 ", ans)
	}

}

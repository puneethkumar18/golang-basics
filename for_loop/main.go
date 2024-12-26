package main

import "fmt"

func main() {
	fmt.Println("We Are Starting For Loop")
	for i := 0; i < 10; i++ {
		fmt.Printf("The Value of i is %d\n", i)
	}
	cnt := 10
	for {
		cnt--
		if cnt%2 == 0 {
			continue
		}
		fmt.Println("The Current Value of cnt is", cnt)

		if cnt == 1 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for index, value := range numbers {
		fmt.Println("The Index is ", index, "the Value is ", value)
	}
}

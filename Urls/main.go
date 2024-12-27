package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Printf("We are starting Web Requests in Go\n")

	myUrl := "https://fakestoreapi.com/products/1"

	fmt.Printf("Type of url is %T\n", myUrl)
	fmt.Println("The Actual Url is : ", myUrl)

	formateUrl, error := url.Parse(myUrl)

	if error != nil {
		fmt.Println("Error Recived :", error)
		return
	} else {

		fmt.Printf("Type of formateUrl is %T\n", formateUrl)
		fmt.Println("The formatted Url is :", formateUrl)
	}

	fmt.Println("The Scheme is : ", formateUrl.Scheme)
	fmt.Println("The Host is : ", formateUrl.Host)
	fmt.Println("The Path is : ", formateUrl.Path)
	fmt.Println("The RawQuery is : ", formateUrl.RawQuery)

}

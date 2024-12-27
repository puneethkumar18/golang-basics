package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("We are starting Web Requests in Go")

	url := "https://fakestoreapi.com/products/1"

	res, error := http.Get(url)

	if error != nil {
		fmt.Println("Error Recived :", error)
		return
	}
	defer res.Body.Close()

	// Reading the content of the Response

	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error In the Conetnt is ", err)
	} else {
		fmt.Println("Content is ", string(content))
	}

}

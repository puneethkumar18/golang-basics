package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Titele    string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CURD.....")

	//var url = string("https://jsonplaceholder.typicode.com/todos/1")

	res, error := http.Get("https://jsonplaceholder.typicode.com/todos/2")

	if error != nil {
		fmt.Println("Some Error Occrred :", error)
		return
	}
	defer res.Body.Close()

	// data, err := ioutil.ReadAll(res.Body)
	// if err == nil {
	// 	fmt.Println("Content of the Response is :", string(data))
	// } else {
	// 	fmt.Println("Some Error in data :", err)
	// }

	var todo Todo

	json.NewDecoder(res.Body).Decode(&todo)

	fmt.Println("Data got from res and structured to TODO is :", todo)

}

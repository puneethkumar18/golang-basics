package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	Id          int
	ProductName string
	Price       float64
	Quantity    int
}

func main() {
	fmt.Println("We are starting Web Requests in Go")

	var product Product
	product = Product{
		Id:          1,
		ProductName: "Laptop",
		Price:       1000.00,
		Quantity:    10,
	}

	fmt.Println("Product Id is : ", product.Id)
	fmt.Println("Product Name is : ", product.ProductName)
	fmt.Println("Product Price is :", product.Price)
	fmt.Println("Product Quantity is the Stock is :", product.Quantity)

	// Type of Encoding the Product Data
	jsonData, error := json.Marshal(product)

	if error != nil {
		fmt.Println("Error in Marshalling the Json Data is : ", error)
	} else {
		fmt.Println("The Json Data is : ", strings.Split(string(jsonData), ","))
	}

	// Type of Decoding the Product Data

	var productData Product
	err := json.Unmarshal(jsonData, &productData)

	if err == nil {
		fmt.Println("ProductData  Name:", productData.ProductName)
	}
}

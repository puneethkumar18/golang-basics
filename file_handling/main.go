package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("We are starting File Handling in Go")

	// // File Handling in Go is done using os package
	// // file, _ := os.Create("file.txt") // crates a file with the name file.txt
	// // defer file.Close()			   // closes the file after the function ends
	// // data := []byte("Hello, This is Puneeth Kumar") // data to be written to the file
	// // data1 := "I am a Software Developer"           // data to be written to the file
	// // os.WriteFile("file.txt", data, 0644)           // writes the data to the file
	// // io.WriteString(file, data1)
	// // io.WriteString(file, data1)
	// // content, _ := os.ReadFile("file.txt")
	// // fmt.Println("Content of the file is:", string(content))

	// file, error := os.Open("example.txt")
	// if error != nil {
	// 	fmt.Println("File Created Successfully")

	// }

	// fmt.Println("File Opened Successfully", error)
	// file.Close()

	// buffer := make([]byte, 1024)

	// for {
	// 	n, error := file.Read(buffer)
	// 	if error == io.EOF {
	// 		break
	// 	}

	// 	if error != nil {
	// 		return
	// 	}

	// 	fmt.Println(string(buffer[:n]))
	// }

	content, error := ioutil.ReadFile("example.txt")
	if error != nil {
		fmt.Println("Error Occuered", error)
		return
	}

	fmt.Println(string(content))
}

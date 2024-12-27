package main

import "fmt"

// Struct Declaration
type Student struct {
	name    string
	age     int
	roolNO  int
	collage string
	section string
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	DoorNo int
	Street string
	City   string
	State  string
}

type Employee struct {
	Person_Details  Person
	Contact_Details Contact
	Address_Details Address
}

func main() {
	fmt.Println("We are starting Structs in Go")

	// // struct initialization
	// var student1 Student

	// // declaration using new keyword
	// student := new(Student)
	// student.name = "Rajesh"
	// student.age = 22
	// student.roolNO = 1
	// student.collage = "NMIT"
	// student.section = "CSE"

	// fmt.Printf("Strundext Address is %p\n", student)

	// fmt.Println("Student  is ", student)

	// student1 = Student{"Puneeth", 22, 1, "NMIT", "CSE"}

	// student2 := Student{"Ramesh", 42, 90, "NIE", "D"}

	// var person1 Person
	// person1.FirstName = "Puneeth"
	// person1.LastName = "Kumar"
	// person1.Age = 22

	// fmt.Println("First Name of the Person is ", person1.FirstName)
	// fmt.Println("Last Name of the Person is ", person1.LastName)
	// fmt.Println("Age of the Person is ", person1.Age)

	// fmt.Println("First Student of the Student Struct is ", student1)
	// fmt.Println("Second Student of the Student Struct is ", student2)

	// declaring a struct with nested struct

	var employee Employee

	var employee1 = new(Employee)

	employee.Person_Details = Person{
		FirstName: "Puneeth",
		LastName:  "Kumar",
		Age:       22,
	}

	employee.Contact_Details.Email = "punin@gmail.com"
	employee.Contact_Details.Phone = "1234567890"
	employee.Address_Details.DoorNo = 123
	employee.Address_Details.Street = "MG Road"
	employee.Address_Details.City = "Bangalore"
	employee.Address_Details.State = "Karnataka"

	fmt.Println("Employee Details are ", employee)
	fmt.Printf("Employee Poiter Address is %p\n", &employee)
	fmt.Printf("Employee1 Poiter Address is %p\n", employee1)
}

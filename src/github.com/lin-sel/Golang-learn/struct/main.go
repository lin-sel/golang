package main

import "fmt"

// Student Details
type Student struct {
	FirstName, LastName, Dept string // Exporting Field
	Age                       int8   // Exporting Field
	address                   string // Non Exporting Field
}

func main() {

	// Creating Student type variable
	var p1 Student

	p1 = Student{
		"Nilesh",
		"Yadav",
		"CSE",
		22,
		"G/8 Jai Nagar Mumbai 101",
	}

	fmt.Println(p1)

	// Assign Pointer to user define type
	point := &p1

	// Accessing Field using pointer
	fmt.Println(point.Age)
}

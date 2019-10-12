package main

import "fmt"

// Student Details
type Student struct {
	FirstName    string
	LastName     string
	Professional string
	RollNumber   int16
	Department   string
	Age          int8
}

func (st *Student) printStudent() string {

	return st.FirstName + " " + st.LastName + " " + string(st.RollNumber) + " " + st.Department + " " + string(st.Age) + " " + st.Professional
}

func main() {

	// Declare Variable of Student type
	var stu Student

	stu = Student{"Nilesh", "Yadav", "Engineer", 4019, "CSE", 22}

	// Calling method
	fmt.Println(stu.printStudent())

}

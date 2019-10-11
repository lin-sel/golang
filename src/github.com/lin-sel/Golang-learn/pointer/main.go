package main

import "fmt"

func main() {

	var vari = 120

	// Regular Pointer Declaration
	var poin *int

	// no value of pointer variable is <nil>
	fmt.Println(poin)

	// Pointer variable also inferred By Compiler so we can omit type.
	var point = &vari

	fmt.Println(point)

	// Accessing Value of Variable pointing By Pointer.
	fmt.Println(*point)

	// Pointer of Pointer
	var poins = &point

	// Accessing Value.
	fmt.Println(poins)
	fmt.Println(*poins)
	fmt.Println(**poins)

	// Creating pointer variable using new keyword
	var p = new(int)

	fmt.Println(p)

	// initialized pointer variable
	*p = 150
	fmt.Println(*p)

	// changing value using pointer
	**poins = 9999

	fmt.Println(vari)

	// Arithmatic Operation
	var t = *p + 1
	fmt.Println(t)

	// Compare pointer
	if *p == *point {
		fmt.Println("Both are Equals.")
	} else {
		fmt.Println("Not Equal")
	}

}

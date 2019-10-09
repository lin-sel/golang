package main

import "fmt"

func main() {

	// Array Declaration
	var arr [4]int

	// Above Array are not just declare it also initialize with 0
	fmt.Println(arr)

	// Output Should be like this
	//[0 0 0 0]

	// Shorthand Array Declaration
	sumApp := [4]int{1, 2, 3, 4}

	fmt.Println(sumApp)
	// Accessing Array with index
	sumApp[0] = 9
	sumApp[1] = 6
	sumApp[2] = 7

	fmt.Println(sumApp)

	// Array Iteration
	for i := 0; i < len(sumApp); i++ {
		fmt.Println(sumApp[i])
	}

	// Array Iteration with range
	for index, value := range sumApp {
		fmt.Printf("value at %d is %d\n", index, value)
	}

	// Array Iteration with range (if we dont care about index or value then we need to use Black identifier '_')
	for _, value := range sumApp {
		fmt.Printf("value is %d\n", value)
	}
}

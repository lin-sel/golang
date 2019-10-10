package main

import "fmt"

func main() {

	// Slice is segment of Array. it provide more power, flexibility and convenience compared to arrays.
	// Unlike Array, we can resized slice.

	// Declaration
	var slice []int

	// Declaration with slice literal
	var slice2 = []int{12, 13, 14, 15, 16}

	fmt.Println(slice2, slice)

	// Shorthand Declaration
	slice3 := []int{19, 18, 17, 16, 15}

	fmt.Println(slice3)

	// Slice Declaration using make function
	len, cap := 4, 10
	slice4 := make([]int, len, cap)

	// Accessing Slice using Indexed
	slice4[0] = 12
	slice4[1] = 13
	slice4[2] = 14
	slice4[3] = 15

	// Iterating slice Using for loop
	for _, value := range slice4 {
		fmt.Println(value)
	}

}

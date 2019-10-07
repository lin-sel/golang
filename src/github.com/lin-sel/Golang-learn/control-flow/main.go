package main

import (
	"fmt"
)

func main() {

	// if else statement
	if 20%2 == 0 {
		fmt.Printf("%d is divisible by %d", 20, 2)
	} else {
		fmt.Println("Not divisible.")
	}

	// shorthand with if-else
	if x := 10; x%2 == 0 {
		fmt.Printf("%d is divisible by %d", x, 2)
	} else {
		fmt.Printf("%d is not divisible by %d", x, 2)
	}

	// In this expression scope of variable will be within if-else chain.

	//---------------------------------------------------------------------

	// Switch

	// (1) Switch with no Expression

	var age int8 = 22

	switch {

	case age > 18:
		{
			fmt.Println("You can eligible for Vote.")
		}

	case age < 18:
		{
			fmt.Println("Not Eligible")
		}

	default:
		{
			fmt.Println("Not Selected")
		}

	}

	// (2) Switch with Shorthand

	switch age := 22; age {

	case 22:
		{
			fmt.Println("You can eligible for Vote.")
		}

	case 18:
		{
			fmt.Println("Not Eligible")
		}

	default:
		{
			fmt.Println("Not Selected")
		}

	}

	// (3) Switch Regular Expression
	ages := 22
	switch ages {

	case 22:
		{
			fmt.Println("You can eligible for Vote.")
		}

	case 18:
		{
			fmt.Println("Not Eligible")
		}

	default:
		{
			fmt.Println("Not Selected")
		}

	}

	// For Loop

	// (1) Regular Expression

	for i := 2; i < 20; i += 2 {
		fmt.Println(i)
	}

	// (2) without initialization statement
	j := 2
	for ; j < 20; j *= 2 {
		fmt.Println(j)
	}

	// (2) without initialization and increment statement
	k := 20
	for k > 0 {
		fmt.Println(k)
		k /= 2
	}
}

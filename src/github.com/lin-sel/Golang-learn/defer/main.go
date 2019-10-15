package main

import "fmt"

func main() {
	behaviour1()
	behaviour2()
	behaviour3()
	behaviour4()
}

// defer statement pushes function into stack. The list of function pushed in stack call after last statement execution of function and before return from function.
// pop function from stack in LIFO Behaviours.

func behaviour1() {
	fmt.Println("Behaviour A")
	defer fmt.Println("Behaviour B")
	fmt.Println("Behaviour C")
}

func behaviour2() {
	defer fmt.Println("Behaviour2 A")
	defer fmt.Println("Behaviour2 B")
	defer fmt.Println("Behaviour2 C")
}

func behaviour3() {
	a := 12
	defer fmt.Println("Behaviour3 ", a)
	a = 20
}

var a int = 12

func behaviour4() {

	defer fmt.Println("Behaviour4 ", a)
	defer swapInt()
}

func swapInt() {
	a = 20
}

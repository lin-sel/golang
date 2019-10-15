package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		fmt.Println(amt)
		time.Sleep(time.Millisecond * amt)
	}
}

/*
This program consists of two goroutines. The first goroutine is implicit and is the main function itself. The second goroutine is created when we call go f(0).
Normally when we invoke a function our program will execute all the statements in a function and then return to the next line following the invocation.
With a goroutine we return immediately to the next line and don't wait for the function to complete. This is why the call to the Scanln function has been included;
without it the program would exit before being given the opportunity to print all the numbers.
*/

// Goroutine is lightweight thread manage by goruntime. we use 'go' keyword to achieve concurrency.
func main() {
	// go f(0)
	// for i := 0; i < 10; i++ {
	// 	go f(i)
	// }
	// var input string
	// fmt.Scanln(&input)
	f(0)
}

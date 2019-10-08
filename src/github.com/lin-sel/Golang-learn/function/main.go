package main

import (
	"errors"
	"fmt"
	"math"
)

// Function is block of code that can be reusable. They improve the readability, maintainability, and testability of your program.
func main() {
	fmt.Println("This is regular expression function.")

	// function calling
	change := changeStockPrice(6.98, 3.98)
	fmt.Println(change)

	// function calling
	changes, percentage := changeStockPricePercentage(6.9898, 3.9898)
	fmt.Println(changes, percentage)

	// function calling
	changes1, percentage1, err := getChangeinStock(6.9898, 3.9898)
	if err == nil {
		fmt.Println(changes1, percentage1)
	} else {
		fmt.Println(err)
	}

	// Blank Identifier
	_, percentage1, err = getChangeinStock(6.9898, 3.9898)
	if err == nil {
		fmt.Println(percentage1)
	} else {
		fmt.Println(err)
	}

}

// function with parameters
func changeStockPrice(prev float64, current float64) (change float32) {
	// type casting
	return float32(current - prev)
}

// function with parameters and multiple return type
func changeStockPricePercentage(prev float64, current float64) /* return value*/ (change, percentage float64) {

	// type casting
	change = (current - prev)
	percentage = change / prev * 100

	// Multiple Return value
	return math.Abs(change), math.Abs(percentage)
}

// function with error return
func getChangeinStock(prev, current float64) (float64, float64, error) {
	if prev == 0 {
		// Custome Error Generating
		err := errors.New("Previous Data should not be Zero")
		return 0, 0, err
	}
	chang := (current - prev)
	percent := chang / prev * 100
	return math.Abs(chang), math.Abs(percent), nil
}

// Sample function for Blank Identifier
func getChangeinStocks(prev, current float64) (float64, float64, error) {
	if prev == 0 {
		// Custome Error Generating
		err := errors.New("Previous Data should not be Zero")
		return 0, 0, err
	}
	chang := (current - prev)
	percent := chang / prev * 100
	return math.Abs(chang), math.Abs(percent), nil
}

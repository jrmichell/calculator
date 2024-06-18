// CLI Calculator written in Go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("The answer is %.2f\n", calculate())

	var response string

	fmt.Println("Do you want to perform another calculation? (y/n)")
	fmt.Scanln(&response)

	for response == "y" {
		fmt.Printf("The answer is %.2f\n", calculate())
		fmt.Println("Do you want to perform another calculation? (y/n)")
		fmt.Scanln(&response)

		if response != "y" {
			fmt.Println("Goodbye!")
			break
		}
	}
}

func calculate() float64 {

	var a, b, result float64

	// Read input from the user
	fmt.Println("Enter the first number:")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number:")
	fmt.Scanln(&b)
	operator := getOperator()

	switch operator {

	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		fmt.Println("Invalid operator.")
	}

	return result
}

func isNumber(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}

func getOperator() string {
	fmt.Println("Enter the operator:")
	var operator string
	fmt.Scanln(&operator)
	return operator
}

func add(a, b float64) float64 {
	if !isNumber(a) || !isNumber(b) {
		fmt.Println("Both arguments must be numbers.")
	}

	return a + b
}

func subtract(a, b float64) float64 {
	if !isNumber(a) || !isNumber(b) {
		fmt.Println("Both arguments must be numbers.")
	}

	return a - b
}

func multiply(a, b float64) float64 {
	if !isNumber(a) || !isNumber(b) {
		fmt.Println("Both arguments must be numbers.")
	}

	return a * b
}

func divide(a, b float64) float64 {
	if !isNumber(a) || !isNumber(b) {
		fmt.Println("Both arguments must be numbers.")
	}

	if b == 0 {
		fmt.Println("Cannot divide by zero.")
	}

	return a / b
}

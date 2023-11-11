package main

import "fmt"

func main() {
	var num1, num2 float64
	var operador string
	fmt.Print("Enter the first number:")
	fmt.Scan(&num1)
	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operador)
	fmt.Print("Enter the second number:")
	fmt.Scan(&num2)
	switch operador {
	case "+":
		fmt.Println("Result:", num1+num2)
	case "-":
		fmt.Println("Result:", num1-num2)
	case "*":
		fmt.Println("Result:", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Println("Result:", num1/num2)
		} else {
			fmt.Println("Error: Cannot divide by zero.")
		}
	default:
		fmt.Println("Invalid operator.")
	}
}

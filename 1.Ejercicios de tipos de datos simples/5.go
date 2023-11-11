package main

import "fmt"

func main() {
	var num1, num2 float64
	fmt.Print("Enter the first decimal number:")
	fmt.Scan(&num1)
	fmt.Print("Enter the second decimal number:")
	fmt.Scan(&num2)
	resultado := num1 + num2
	fmt.Println("The result of the sum is:", resultado)
}

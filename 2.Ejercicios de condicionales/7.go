/*
Que pida 3 números y los muestre en pantalla de menor a mayor.
*/
package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter the third number: ")
	fmt.Scan(&num3)
	numbers := []int{num1, num2, num3}
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	fmt.Println("Numbers ordered from least to greatest:", numbers)
}

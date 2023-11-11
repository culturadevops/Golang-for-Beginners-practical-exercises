/*Que pida tres números y detecte si se han introducido en orden decreciente.*/
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
	if num1 > num2 && num2 > num3 {
		fmt.Println("The numbers are in descending order.")
	} else {
		fmt.Println("The numbers are not in descending order.")
	}
}

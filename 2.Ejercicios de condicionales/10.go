/*
Que pida un número y diga si es mayor de 100.
*/
package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if num > 100 {
		fmt.Println("The number is greater than 100.")
	} else {
		fmt.Println("The number is not greater than 100.")
	}
}

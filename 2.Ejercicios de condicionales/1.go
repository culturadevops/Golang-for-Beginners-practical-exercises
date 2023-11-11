/*
Escribir un programa que pregunte al usuario su edad y muestre por pantalla
si es mayor de edad o no
*/
package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You are of legal age.")
	} else {
		fmt.Println("You are a minor.")
	}
}

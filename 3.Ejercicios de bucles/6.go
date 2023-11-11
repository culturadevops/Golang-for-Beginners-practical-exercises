/*Escribir un programa que pregunte al usuario su edad y muestre por pantalla todos los años que ha cumplido (desde 1 hasta su edad).*/

package main

import "fmt"

func main() {
	var int age
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	for i := 1; i <= age; i++ {
		fmt.Println("Has turned", i, "years old.")
	}
}

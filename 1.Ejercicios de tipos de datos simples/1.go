package main

import "fmt"

// Programa que pide dos enteros y los multiplica:
func main() {
	var num1, num2 int
	fmt.Print("Ingrese el primer entero: ")
	fmt.Scan(&num1)
	fmt.Print("Ingrese el segundo entero: ")
	fmt.Scan(&num2)
	resultado := num1 * num2
	fmt.Println("El resultado de la multiplicación es:", resultado)
}

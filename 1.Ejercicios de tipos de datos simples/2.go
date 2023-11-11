package main

import "fmt"

// Programa que pide un decimal y lo divide:
func main() {
	var num float64
	fmt.Print("Ingrese un número decimal: ")
	fmt.Scan(&num)
	fmt.Println("El número dividido por 2 es:", num/2)
}

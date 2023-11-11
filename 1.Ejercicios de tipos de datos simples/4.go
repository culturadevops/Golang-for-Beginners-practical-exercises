package main

import "fmt"

// Programa que pide al usuario tres valores (int, float, y string) y los muestra:
func main() {
	var numInt int
	var numFloat float64
	var texto string

	fmt.Print("Ingrese un número entero: ")
	fmt.Scan(&numInt)
	fmt.Print("Ingrese un número decimal: ")
	fmt.Scan(&numFloat)
	fmt.Print("Ingrese un texto: ")
	fmt.Scan(&texto)

	fmt.Println("El número entero ingresado es:", numInt)
	fmt.Println("El número decimal ingresado es:", numFloat)
	fmt.Println("El texto ingresado es:", texto)
}

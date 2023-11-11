package main

import "fmt"

// Programa que pide un string y lo muestra:
func main() {
	var texto string
	fmt.Print("Ingrese un texto: ")
	fmt.Scan(&texto)
	fmt.Println("El texto ingresado es:", texto)
}

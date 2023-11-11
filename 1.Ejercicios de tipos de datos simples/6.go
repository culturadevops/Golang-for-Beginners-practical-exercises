package main

import "fmt"

func main() {
	var num int
	fmt.Print("Ingrese un número entero: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("El número es par.")
	} else {
		fmt.Println("El número es impar.")
	}
}

// Programa que pregunta al usuario por un número y valida si es par o no:

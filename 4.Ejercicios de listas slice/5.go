/*Escribir un programa que almacene en una lista los números del 1 al 10 y los muestre por pantalla en orden inverso separados por comas.*/
package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Números en orden inverso:")
	for i := len(numeros) - 1; i >= 0; i-- {
		fmt.Printf("%d", numeros[i])

		if i != 0 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

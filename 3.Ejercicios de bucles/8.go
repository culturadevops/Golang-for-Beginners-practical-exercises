/*
Escribir un programa que pida al usuario un número entero positivo y
muestre por pantalla la cuenta atrás desde ese número
hasta cero separados por comas.
*/
package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&num)

	fmt.Print("Countdown:")
	for i := num; i >= 0; i-- {
		fmt.Print(" ", i)
	}
	fmt.Println()
}

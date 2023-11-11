/*Escribir un programa que pida al usuario un número entero positivo y
muestre por pantalla todos los números impares desde 1 hasta ese número separados por comas.*/

package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&num)
	fmt.Print("Odd numbers:")
	for i := 1; i <= num; i += 2 {
		fmt.Print(" ", i)
	}
	fmt.Println()
}

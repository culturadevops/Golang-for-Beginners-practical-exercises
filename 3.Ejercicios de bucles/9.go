/*
Escribir un programa que pida al usuario un número entero y muestre por pantalla un triángulo
rectángulo como el de más abajo, de altura el número introducido.
*/
package main

import (
	"fmt"
)

func main() {
	var height int
	fmt.Print("Enter the height of the triangle: ")
	fmt.Scan(&height)
	for i := 1; i <= height; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

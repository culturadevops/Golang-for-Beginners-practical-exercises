/*Escribir un programa que muestre por pantalla la tabla de multiplicar del 1 al 10.*/
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Multiplication table", i, ":")
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "x", j, "=", i*j)
		}
		fmt.Println()
	}
}

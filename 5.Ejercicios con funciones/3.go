/*Escribir una función que reciba un número entero positivo y devuelva su factorial.
 */
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	numero := 5
	fmt.Printf("El factorial de %d es %d\n", numero, factorial(numero))
}

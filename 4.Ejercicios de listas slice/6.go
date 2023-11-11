/*Escriba un programa que recorra una lista de numero y la copie en otra de forma invertida.*/
package main

import (
	"fmt"
)

func main() {
	numerosOriginal := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numerosInvertidos := make([]int, len(numerosOriginal))
	for i := 0; i < len(numerosOriginal); i++ {
		numerosInvertidos[i] = numerosOriginal[len(numerosOriginal)-1-i]
	}
	fmt.Println("Números originales:", numerosOriginal)
	fmt.Println("Números invertidos:", numerosInvertidos)
}

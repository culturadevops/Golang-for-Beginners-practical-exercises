/*Escribir una función que reciba una muestra de números en una lista y devuelva otra lista con sus cuadrados.
 */
package main

import (
	"fmt"
)

func cuadrados(numeros []int) []int {
	cuadrados := []int{}
	for _, num := range numeros {
		cuadrado := num * num
		cuadrados = append(cuadrados, cuadrado)
	}
	return cuadrados
}
func main() {
	muestra := []int{1, 2, 3, 4, 5}
	fmt.Println("Muestra original:", muestra)
	fmt.Println("Cuadrados:", cuadrados(muestra))
}

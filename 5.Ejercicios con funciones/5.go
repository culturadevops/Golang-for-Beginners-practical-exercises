/*Escribir una función que reciba una muestra de números en una lista y devuelva su media.
 */
package main

import (
	"fmt"
)

func media(numeros []float64) float64 {
	suma := 0.0
	for _, num := range numeros {
		suma += num
	}
	return suma / float64(len(numeros))
}
func main() {
	muestra := []float64{5, 7, 3, 10, 8}
	fmt.Printf("Media de la muestra: %.2f\n", media(muestra))
}

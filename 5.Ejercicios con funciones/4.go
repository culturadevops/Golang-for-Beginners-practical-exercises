/*Escribir una función que calcule el área de un círculo y otra que calcule el volumen de un cilindro usando la primera función.
 */
package main

import (
	"fmt"
	"math"
)

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}
func volumenCilindro(radio, altura float64) float64 {
	return areaCirculo(radio) * altura
}
func main() {
	radio := 2.0
	altura := 5.0
	fmt.Printf("Área del círculo: %.2f\n", areaCirculo(radio))
	fmt.Printf("Volumen del cilindro: %.2f\n", volumenCilindro(radio, altura))
}

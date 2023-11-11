/*Añadir y remover ítems al principio y al final de un arreglo
 */
package main

import (
	"fmt"
)

func main() {
	arreglo := []int{1, 2, 3, 4, 5}
	// Añadir al principio
	arreglo = append([]int{0}, arreglo...)
	// Añadir al final
	arreglo = append(arreglo, 6)
	fmt.Println("Arreglo después de agregar al principio y al final:", arreglo)
	// Remover del principio
	arreglo = arreglo[1:]
	// Remover del final
	arreglo = arreglo[:len(arreglo)-1]
	fmt.Println("Arreglo después de remover al principio y al final:", arreglo)
}

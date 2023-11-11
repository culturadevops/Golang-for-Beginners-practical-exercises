/*Escribir un programa que pregunte al usuario los números ganadores de la lotería primitiva, los almacene en una lista y los muestre por pantalla ordenados de menor a mayor.
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	var ganadores []int
	for i := 1; i <= 6; i++ {
		var numero int
		fmt.Printf("Ingrese el número ganador %d: ", i)
		fmt.Scan(&numero)
		ganadores = append(ganadores, numero)
	}
	sort.Ints(ganadores)
	fmt.Println("Números ganadores ordenados:")
	fmt.Println(ganadores)
}

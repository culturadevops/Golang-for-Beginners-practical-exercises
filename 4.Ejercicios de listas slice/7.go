/*Escribir un programa que almacene las asignaturas de un curso (por ejemplo Matemáticas, Física, Química, Historia y Lengua) en una lista, pregunte al usuario la nota que ha sacado en cada asignatura y elimine de la lista las asignaturas aprobadas. Al final el programa debe mostrar por pantalla las asignaturas que el usuario tiene que repetir.*/
package main

import (
	"fmt"
)

func main() {
	asignaturas := []string{"Matemáticas", "Física", "Química", "Historia", "Lengua"}
	notas := make(map[string]float64)
	asignaturasARepetir := []string{}

	for _, asignatura := range asignaturas {
		var nota float64
		fmt.Printf("Ingrese la nota de %s: ", asignatura)
		fmt.Scan(&nota)
		notas[asignatura] = nota

		if nota < 6 {
			asignaturasARepetir = append(asignaturasARepetir, asignatura)
		}
	}

	fmt.Println("Asignaturas a repetir:")
	for _, asignatura := range asignaturasARepetir {
		fmt.Println(asignatura)
	}
}

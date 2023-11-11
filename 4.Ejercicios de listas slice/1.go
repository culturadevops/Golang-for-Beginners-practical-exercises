/*Escribir un programa que almacene las asignaturas de un curso (por ejemplo Matemáticas, Física, Química, Historia y Lengua) en una lista y la muestre por pantalla.*/
package main

import (
	"fmt"
)

func main() {
	asignaturas := []string{"Matemáticas", "Física", "Química", "Historia", "Lengua"}
	fmt.Println("Asignaturas del curso:")
	for _, asignatura := range asignaturas {
		fmt.Println(asignatura)
	}
}

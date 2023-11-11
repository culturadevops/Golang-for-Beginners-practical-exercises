/*Escribir un programa que almacene las asignaturas de un curso (por ejemplo Matemáticas, Física, Química, Historia y Lengua) en una lista, pregunte al usuario la nota que ha sacado en cada asignatura, y después las muestre por pantalla con el mensaje En <asignatura> has sacado <nota> donde <asignatura> es cada una des las asignaturas de la lista y <nota> cada una de las correspondientes notas introducidas por el usuario. adicional a esto debe decir si ha pasado la asignatura si la nota es mayor a 6*/
package main

import (
	"fmt"
)

func main() {
	asignaturas := []string{"Matemáticas", "Física", "Química", "Historia", "Lengua"}
	notas := make(map[string]float64)
	for _, asignatura := range asignaturas {
		var nota float64
		fmt.Printf("Ingrese la nota de %s: ", asignatura)
		fmt.Scan(&nota)
		notas[asignatura] = nota
	}
	fmt.Println("Notas obtenidas:")
	for asignatura, nota := range notas {
		fmt.Printf("En %s has sacado %.2f", asignatura, nota)
		if nota >= 6 {
			fmt.Println(" - Has aprobado.")
		} else {
			fmt.Println(" - Has reprobado.")
		}
	}
}

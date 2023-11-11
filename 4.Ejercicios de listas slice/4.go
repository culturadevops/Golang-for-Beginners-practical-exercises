/*Escribir un programa en el que se pregunte al usuario por una frase y una letra, y muestre por pantalla el número de veces que aparece la letra en la frase.*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var frase string
	var letra string

	fmt.Print("Ingrese una frase: ")
	fmt.Scan(&frase)

	fmt.Print("Ingrese una letra: ")
	fmt.Scan(&letra)

	contador := strings.Count(frase, letra)
	fmt.Printf("La letra '%s' aparece %d veces en la frase.\n", letra, contador)
}

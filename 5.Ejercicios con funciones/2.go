/*Escribir una función a la que se le pase una cadena <nombre> y muestre por pantalla el saludo ¡hola <nombre>!.
 */
package main

import (
	"fmt"
)

func saludar(nombre string) {
	fmt.Printf("¡Hola %s!\n", nombre)
}

func main() {
	nombre := "María"
	saludar(nombre)
}

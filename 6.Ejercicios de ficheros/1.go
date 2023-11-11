/*Haz un programa que lea un archivo y lo imprima en pantalla
 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	nombreArchivo := "archivo.txt"

	contenido, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		log.Fatalf("Error al leer el archivo: %v", err)
	}

	fmt.Println("Contenido del archivo:")
	fmt.Println(string(contenido))
}

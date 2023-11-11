/*Haz un programa que pida un texto al usuario y luego lo guarde en un archivo
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	nombreArchivo := "archivo_usuario.txt"

	fmt.Print("Ingrese un texto: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()

	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		log.Fatalf("Error al crear el archivo: %v", err)
	}
	defer archivo.Close()

	_, err = archivo.WriteString(texto)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}

	log.Println("Texto guardado exitosamente en el archivo.")
}

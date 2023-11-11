/*Escribe un programa que lea un archivo y lo copie en otro archivo
 */
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	nombreArchivoOrigen := "archivo_origen.txt"
	nombreArchivoDestino := "archivo_destino.txt"

	archivoOrigen, err := os.Open(nombreArchivoOrigen)
	if err != nil {
		log.Fatalf("Error al abrir el archivo de origen: %v", err)
	}
	defer archivoOrigen.Close()

	archivoDestino, err := os.Create(nombreArchivoDestino)
	if err != nil {
		log.Fatalf("Error al crear el archivo de destino: %v", err)
	}
	defer archivoDestino.Close()

	_, err = io.Copy(archivoDestino, archivoOrigen)
	if err != nil {
		log.Fatalf("Error al copiar el archivo: %v", err)
	}

	log.Println("Archivo copiado exitosamente.")
}

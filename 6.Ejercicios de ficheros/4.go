/*Haz un programa que abra un archivo de texto y lo convierta en un archivo HTML.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	nombreArchivoTexto := "archivo_texto.txt"
	nombreArchivoHTML := "archivo_html.html"

	contenidoTexto, err := ioutil.ReadFile(nombreArchivoTexto)
	if err != nil {
		log.Fatalf("Error al leer el archivo de texto: %v", err)
	}

	contenidoHTML := fmt.Sprintf("<html><body><p>%s</p></body></html>", string(contenidoTexto))

	archivoHTML, err := os.Create(nombreArchivoHTML)
	if err != nil {
		log.Fatalf("Error al crear el archivo HTML: %v", err)
	}
	defer archivoHTML.Close()

	_, err = archivoHTML.WriteString(contenidoHTML)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo HTML: %v", err)
	}

	log.Println("Archivo HTML generado exitosamente.")
}

/*Leer un archivo con una tabla HTML, remover las etiquetas y guardar la información en otro archivo:
 */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	nombreArchivoHTML := "tabla_html.html"
	nombreArchivoDatos := "datos_tabla.txt"
	archivoHTML, err := os.Open(nombreArchivoHTML)
	if err != nil {
		log.Fatalf("Error al abrir el archivo HTML: %v", err)
	}
	defer archivoHTML.Close()
	archivoDatos, err := os.Create(nombreArchivoDatos)
	if err != nil {
		log.Fatalf("Error al crear el archivo de datos: %v", err)
	}
	defer archivoDatos.Close()
	scanner := bufio.NewScanner(archivoHTML)
	for scanner.Scan() {
		linea := scanner.Text()
		lineaSinEtiquetas := quitarEtiquetasHTML(linea)
		fmt.Fprintln(archivoDatos, lineaSinEtiquetas)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error al leer el archivo HTML: %v", err)
	}
	log.Println("Datos extraídos de la tabla HTML y guardados en el archivo.")
}
func quitarEtiquetasHTML(texto string) string {
	return strings.TrimSpace(regExpEtiquetas.ReplaceAllString(texto, ""))
}

// Regexp para encontrar las etiquetas HTML
var regExpEtiquetas = regexp.MustCompile("<[^>]*>")

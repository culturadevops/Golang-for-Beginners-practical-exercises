/*Extrae palabras en mayúsculas de un archivo.
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extraerMayusculas(archivo string) []string {
	archivoEntrada, err := os.Open(archivo)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return nil
	}
	defer archivoEntrada.Close()
	mayusculas := []string{}
	scanner := bufio.NewScanner(archivoEntrada)
	for scanner.Scan() {
		linea := scanner.Text()
		palabras := strings.Fields(linea)

		for _, palabra := range palabras {
			if palabra == strings.ToUpper(palabra) {
				mayusculas = append(mayusculas, palabra)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return nil
	}
	return mayusculas
}

func main() {
	nombreArchivo := "archivo.txt"
	palabrasMayusculas := extraerMayusculas(nombreArchivo)
	if palabrasMayusculas != nil {
		fmt.Println("Palabras en mayúsculas encontradas:")
		for _, palabra := range palabrasMayusculas {
			fmt.Println(palabra)
		}
	}
}

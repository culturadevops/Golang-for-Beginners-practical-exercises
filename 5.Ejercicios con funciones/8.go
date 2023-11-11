/*Haz un programa que obtenga la hora y fecha del sistema y lo convierta a diferentes formatos como “DD/MM/AAAA” o “<dia> del <mes > de <año>” etc.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fechaActual := time.Now()
	// Formato "DD/MM/AAAA"
	fmt.Println("Formato DD/MM/AAAA:", fechaActual.Format("02/01/2006"))
	// Formato "<día> del <mes> de <año>"
	fmt.Println("Formato <día> del <mes> de <año>:", fechaActual.Format("2 del 1 de 2006"))
	// Formato "AAAA-MM-DD HH:mm:ss"
	fmt.Println("Formato AAAA-MM-DD HH:mm:ss:", fechaActual.Format("2006-01-02 15:04:05"))
}

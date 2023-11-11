/*Escribir una función que convierta un número decimal en binario y otra que convierta un número binario en decimal.
 */
package main

import (
	"fmt"
	"strconv"
)

func decimalABinario(decimal int) string {
	return strconv.FormatInt(int64(decimal), 2)
}
func binarioADecimal(binario string) (int, error) {
	return strconv.ParseInt(binario, 2, 64)
}
func main() {
	numeroDecimal := 25
	fmt.Printf("Decimal %d en binario: %s\n", numeroDecimal, decimalABinario(numeroDecimal))
	numeroBinario := "11001"
	numeroDecimal, err := binarioADecimal(numeroBinario)
	if err != nil {
		fmt.Println("Error al convertir el número binario a decimal:", err)
	} else {
		fmt.Printf("Binario %s en decimal: %d\n", numeroBinario, numeroDecimal)
	}
}

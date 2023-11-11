/*
Escribir un programa que pida al usuario una palabra y luego muestre por pantalla una a una las letras de la palabra introducida empezando por la última.
*/
package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scan(&word)
	fmt.Println("Letters of the word in reverse order:")
	for i := len(word) - 1; i >= 0; i-- {
		fmt.Println(string(word[i]))
	}
}

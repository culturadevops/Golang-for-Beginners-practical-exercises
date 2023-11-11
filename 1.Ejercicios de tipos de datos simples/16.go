/*
	Programa que pide al usuario 5 palabras, las guarda en un arreglo,

las ordena y muestra el resultado:
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	words := make([]string, 5)
	for i := 0; i < len(words); i++ {
		fmt.Print("Enter a word: ")
		fmt.Scan(&words[i])
	}
	sort.Strings(words)
	fmt.Println("Ordered words:", words)
}

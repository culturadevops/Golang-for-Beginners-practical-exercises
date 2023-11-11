// Programa que concatena strings sin usar el operador +:
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	text1 := "Hello"
	text2 := "World"

	buffer.WriteString(text1)
	buffer.WriteString(" ")
	buffer.WriteString(text2)

	fmt.Println(buffer.String())
}

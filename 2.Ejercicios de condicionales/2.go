/*
Escribir un programa que almacene la cadena de caracteres contraseña en una variable,
pregunte al usuario por la contraseña e imprima por pantalla si la contraseña
introducida por el usuario coincide con la guardada en la variable sin tener
en cuenta mayúsculas y minúsculas
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	password := "MySecurePassword"
	var inputPassword string
	fmt.Print("Enter your password: ")
	fmt.Scan(&inputPassword)

	if strings.EqualFold(password, inputPassword) {
		fmt.Println("Password correct.")
	} else {
		fmt.Println("Incorrect password.")
	}
}

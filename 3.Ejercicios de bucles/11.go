/*
Escribir un programa que almacene la cadena de caracteres contraseña en una variable,
pregunte al usuario por la contraseña hasta que introduzca la contraseña correcta.
*/
package main

import "fmt"

func main() {
	password := "MySecurePassword"
	for {
		var inputPassword string
		fmt.Print("Enter your password: ")
		fmt.Scan(&inputPassword)
		if inputPassword == password {
			fmt.Println("Password correct.")
			break
		} else {
			fmt.Println("Incorrect password. Please try again.")
		}
	}
}

/*Que pida una letra y detecte si es una vocal.*/
package main

import "fmt"

func main() {
	var letter string
	fmt.Print("Enter a letter: ")
	fmt.Scan(&letter)
	switch letter {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("It's a vowel.")
	default:
		fmt.Println("Not a vowel.")
	}
}

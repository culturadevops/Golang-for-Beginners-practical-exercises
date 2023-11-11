/*
Que sólo permita introducir los caracteres S y N.
*/
package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter 'Y' or 'N': ")
	fmt.Scan(&input)
	if input == "Y" || input == "y" {
		fmt.Println("You have entered 'Y'")
	} else if input == "N" || input == "n" {
		fmt.Println("You have entered 'N'")
	} else {
		fmt.Println("Invalid character. Only 'Y' or 'N' allowed.")
	}
}

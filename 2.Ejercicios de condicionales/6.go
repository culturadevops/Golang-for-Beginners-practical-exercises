/*
Que pida un número del 1 al 12 y diga el nombre del mes correspondiente.
(use case y luego else if), ¿cual le parece mejor? ¿con cuál te sientes más cómodo?
*/
package main

import "fmt"

func main() {
	var month int
	fmt.Print("Enter a number from 1 to 12: ")
	fmt.Scan(&month)
	if month >= 1 && month <= 12 {
		switch month {
		case 1:
			fmt.Println("January")
		case 2:
			fmt.Println("February")
		case 3:
			fmt.Println("March")
		case 4:
			fmt.Println("April")
		case 5:
			fmt.Println("May")
		case 6:
			fmt.Println("June")
		case 7:
			fmt.Println("July")
		case 8:
			fmt.Println("August")
		case 9:
			fmt.Println("September")
		case 10:
			fmt.Println("October")
		case 11:
			fmt.Println("November")
		case 12:
			fmt.Println("December")
		}
	} else {
		fmt.Println("Invalid number. Must be 1 to 12.")
	}
}

/*
Que pida un número del 1 al 7 y diga el día de la semana correspondiente.
(use case y luego else if),
¿cual le parece mejor?
¿con cuál te sientes más cómodo?
*/
package main

import "fmt"

func main() {
	var day int
	fmt.Print("Enter a number from 1 to 7: ")
	fmt.Scan(&day)
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid number. Must be 1 to 7.")
	}
}

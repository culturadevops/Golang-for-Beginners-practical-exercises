package main

import "fmt"

func main() {
	var initialinvestment, annualinterest float64
	var nyears int

	fmt.Print("Enter the amount to invest: ")
	fmt.Scan(&initialinvestment)
	fmt.Print("Enter the annual interest (%): ")
	fmt.Scan(&annualinterest)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&nyears)

	capitalObtained := initialinvestment * (1 + (annualinterest/100)*float64(nyears))
	fmt.Println("The capital obtained in the investment is:", capitalObtained)
}

// Programa que pregunta al usuario una cantidad a invertir, el interés anual, y el número de años, y muestra el capital obtenido en la inversión:

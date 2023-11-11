package main

import "fmt"

func main() {
	var hoursWorked float64
	var costPerHour float64
	fmt.Print("Enter the number of hours worked:")
	fmt.Scan(&hoursWorked)
	fmt.Print("Enter the cost per hour:")
	fmt.Scan(&costPerHour)
	pay := hoursWorked * costPerHour
	fmt.Println("The corresponding pay is:", pay)
}

// Programa que pregunta al usuario por el número de horas trabajadas y el coste por hora, y muestra la pay correspondiente:

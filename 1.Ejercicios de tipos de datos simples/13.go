// Programa que lee el número de barras vendidas que no son del día y muestra el precio, descuento y coste total:
package main

import "fmt"

func main() {
	var barsNoDay int
	fmt.Print("Enter the number of bars sold today: ")
	fmt.Scan(&barsNoDay)
	normalprice := 1.0
	discount := 0.6
	finalprice := normalprice * (1 - discount)
	totalcost := float64(barsNoDay) * finalprice
	fmt.Printf("Regular price of a loaf of bread: %.2f coins\n", normalprice)
	fmt.Printf("Discount for not being fresh: %.2f coins\n", normalprice-finalprice)
	fmt.Printf("Total final cost: %.2f coins\n", totalcost)
}

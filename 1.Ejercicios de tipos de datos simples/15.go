/* Programa que pide al usuario 10 números y los guarda en un arreglo de enteros:*/
package main

import "fmt"

func main() {
	var numbers [10]int
	for i := 0; i < len(numbers); i++ {
		fmt.Print("Enter a number: ")
		fmt.Scan(&numbers[i])
	}
	fmt.Println("Numbers entered:", numbers)
}

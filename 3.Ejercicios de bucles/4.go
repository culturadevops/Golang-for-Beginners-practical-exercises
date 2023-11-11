/*Pida 10 números al usuario y guardarlos en un arreglo luego imprime uno a uno cada dato introducido de esta manera.*/
package main

import "fmt"

func main() {
	var numbers [10]int
	for i := 0; i < len(numbers); i++ {
		fmt.Print("Enter the number ", i+1, ": ")
		fmt.Scan(&numbers[i])
	}
	fmt.Println("Numbers entered:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

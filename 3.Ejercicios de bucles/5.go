/*Acepta series de números desde el teclado y ordenarlos ascendente o descendentemente (¿puedes hacer las dos opciones?).*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Enter the number of numbers: ")
	fmt.Scan(&n)
	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Enter the number ", i+1, ": ")
		fmt.Scan(&numbers[i])
	}
	fmt.Println("Numbers entered:", numbers)
	// Ascending order
	sort.Ints(numbers)
	fmt.Println("Ascending order:", numbers)
	// Descending order
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	fmt.Println("Descending order:", numbers)
}

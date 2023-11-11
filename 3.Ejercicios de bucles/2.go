/*Muestra la serie de Fibonacci.*/
package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
func main() {
	var num int
	fmt.Print("Enter the number of elements of the Fibonacci series: ")
	fmt.Scan(&num)
	fmt.Println("Fibonacci series:")
	for i := 0; i < num; i++ {
		fmt.Print(fibonacci(i), " ")
	}
	fmt.Println()
}

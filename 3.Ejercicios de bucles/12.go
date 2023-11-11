/* Escribir un programa que pida al usuario 10 números enteros y muestre por pantalla todos los números señalando a los que son primos.
 */
package main

import "fmt"

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var numbers [10]int
	for i := 0; i < len(numbers); i++ {
		fmt.Print("Enter the number ", i+1, ": ")
		fmt.Scan(&numbers[i])
	}
	fmt.Println("Numbers entered:")
	for _, num := range numbers {
		if isPrime(num) {
			fmt.Println(num, "is prime.")
		} else {
			fmt.Println(num, "not prime.")
		}
	}
}

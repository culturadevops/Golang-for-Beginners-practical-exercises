/*Que pida un número del 1 al 5 y diga si es primo o no.*/
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
	var num int
	fmt.Print("Enter a number from 1 to 5: ")
	fmt.Scan(&num)
	if num >= 1 && num <= 5 {
		if isPrime(num) {
			fmt.Println("The number is prime.")
		} else {
			fmt.Println("The number is not prime.")
		}
	} else {
		fmt.Println("Invalid number. Must be 1 to 5.")
	}
}

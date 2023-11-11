/*Encuentra el máximo y el mínimo en una lista de números.*/
package main

import (
	"fmt"
)

func main() {
	var n, num, max, min int
	fmt.Print("Enter the number of numbers: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("Enter the number ", i, ": ")
		fmt.Scan(&num)
		if i == 1 {
			max = num
			min = num
		} else {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
	}
	fmt.Println("The maximum is:", max)
	fmt.Println("The minimum is:", min)
}

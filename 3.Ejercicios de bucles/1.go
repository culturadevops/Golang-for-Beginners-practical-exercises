/*
Muestra series de números en un loop infinito.
El programa debería salir si el usuario presiona una tecla específica.
(por ejemplo, la tecla 'q')
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Press the 'q' key to exit.")
	for {
		// Clear the screen before displaying the next number.
		clearScreen()
		// Show the series of numbers.
		for i := 1; ; i++ {
			fmt.Println(i)
			// Read the character entered by the user.
			var input string
			fmt.Scan(&input)
			// Exit the loop if the user presses the 'q' key.
			if input == "q" {
				os.Exit(0)
			}
		}
	}
}

// Function to clear the screen.
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

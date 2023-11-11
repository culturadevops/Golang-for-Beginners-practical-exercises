/*Que muestre un menú que contemple las opciones “Archivo”, “Buscar” y “Salir”, en caso de que no se introduzca una opción correcta se notificará por pantalla.*/
package main

import "fmt"

func main() {
	var option int
	fmt.Println("Menu:")
	fmt.Println("1. File")
	fmt.Println("2. Search")
	fmt.Println("3. Exit")
	fmt.Print("Enter an option: ")
	fmt.Scan(&option)
	switch option {
	case 1:
		fmt.Println("You have selected the File option.")
	case 2:
		fmt.Println("You have selected the Search option.")
	case 3:
		fmt.Println("You have selected the Exit option.")
	default:
		fmt.Println("Invalid option.")
	}
}

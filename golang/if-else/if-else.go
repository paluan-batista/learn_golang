package main

import "fmt"

func main() {

	fmt.Println("Estruturas de Controle")

	// IF e ELSE

	numero := 10

	if numero > 15 {
		fmt.Print("é maior")
	} else {
		fmt.Println("é menor")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("é maior que zero")
	} else {
		fmt.Println("é menor que zero")
	}

}

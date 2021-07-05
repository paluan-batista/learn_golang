package main

import (
	"fmt"
)

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total

}

func main() {
	fmt.Println(" FUNÇOES VARIÁTICAS")

	result := soma(2, 50, 300, 800, 333, 666)

	fmt.Println(result)

}

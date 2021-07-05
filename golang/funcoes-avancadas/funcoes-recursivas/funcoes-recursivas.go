package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	} else {
		return fibonacci(posicao-2) + fibonacci(posicao-1)
	}

}

func main() {

	fmt.Println("FUNÇÕES RECURSIVAS")

	posicao := uint(55)
	fmt.Println(fibonacci(posicao))
}

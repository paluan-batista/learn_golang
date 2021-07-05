package main

import "fmt"

const texto string = "E estou gostando muito"

func main() {
	fmt.Println("FUNÇOES ANONIMAS")

	func() {
		fmt.Println("ola mundo")
	}()

	func(text string) {
		fmt.Println(text)
	}("Estou aprendendo go")

	retorno := func() string {
		return texto
	}()

	fmt.Println(retorno)
}

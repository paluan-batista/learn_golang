package main

import "fmt"

func closure() func() {
	texto := "dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main() {
	fmt.Println("FUNÇÃO CLOSURE")
	texto := "dentro da função main"

	fmt.Println(texto)
	funcaoNova := closure()

	funcaoNova()
}

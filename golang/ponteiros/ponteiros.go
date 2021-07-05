package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	variavel1 := 10
	variavel2 := variavel1

	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO E UMA REFERENCIA DE MEMORIA

	var variavel3 int = 100
	var ponteiro *int = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //DESREFERENCIAÇÂO

}

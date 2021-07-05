package main

import (
	"errors"
	"fmt"
)

func main() {

	// NUMEROS INTEIROS
	var numero int64 = -100000000000000000
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	// int32 = rune
	var numero3 rune = 1215151
	fmt.Println(numero3)

	// uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	// NUMEROS REAIS
	var numeroReal float32 = 123.456
	fmt.Println(numeroReal)

	// STRING
	var test string = "test"
	test2 := "test2"

	fmt.Println(test)
	fmt.Println(test2)

	//BOOLEANO
	var booleana1 bool = true
	booleana2 := false
	fmt.Println(booleana1)
	fmt.Println(booleana2)

	//ERRO
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("erro interno")
	fmt.Println(erro2)

}

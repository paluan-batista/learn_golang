package main

import "fmt"

func main() {
	var test string = "variavel1"
	fmt.Println(test)

	variavel2 := "variavel2"
	fmt.Println(variavel2)

	var (
		variavel3 string = "paluan"
		variavel4 string = "batista"
	)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "batatais", "sp"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "constante1"
	fmt.Println(constante1)

	//invertendo valores de variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5)
	fmt.Println(variavel6)
}

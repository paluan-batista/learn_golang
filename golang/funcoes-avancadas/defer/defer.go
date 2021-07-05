package main

import "fmt"

// func funcao1() {
// 	fmt.Println("executando a func 1")
// }

// func funcao2() {
// 	fmt.Println("executando a func 2")
// }

func calculaMedia(n1, n2 float32) bool {
	defer fmt.Println("Calculando média...")

	media := (n1 + n2) / 2

	return media >= 6

}
func main() {
	fmt.Println("DEFER")

	// defer funcao1()
	// DEFER = ADIAR

	// funcao2()

	var resultado string

	if calculaMedia(5, 4) {
		resultado = "aprovado"
	} else {
		resultado = "reprovado"
	}

	fmt.Println(resultado)

}

package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}
func calculaMedia(n1, n2 float32) bool {
	defer recuperarExecucao()
	defer fmt.Println("Calculando média...")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é examanntente 6!")

}

func main() {
	fmt.Println("PANIC e RECOVER")
	fmt.Println(calculaMedia(6, 6))
	fmt.Println("FIM DA EXECUÇAO")

}

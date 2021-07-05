package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Estou ententendendo")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("paluan@test.com")
	fmt.Println(erro)

}

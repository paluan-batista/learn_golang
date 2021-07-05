package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso string
}

func main() {

	p1 := pessoa{"Paluan", "Batista", 29}
	fmt.Println(p1)

	estudante1 := estudante{p1, "ads"}

	fmt.Println("Estudante :", estudante1)
	fmt.Println(estudante1.curso)

}

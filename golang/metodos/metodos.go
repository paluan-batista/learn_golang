package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Println("salvando de mentirinha...")
}

func main() {
	fmt.Println("METODOS")
	usuario1 := usuario{"Paluan", 29}
	fmt.Println(usuario1)
	usuario1.salvar()

}

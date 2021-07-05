package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {

	var u usuario

	u.nome = "Paluan"
	u.idade = 29
	fmt.Println(u)

	u2 := usuario{"vanessa", 26}
	fmt.Println(u2)

	u3 := usuario{nome: "eu"}
	fmt.Println(u3)
}

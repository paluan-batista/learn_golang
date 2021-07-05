package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("LOOPS")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando ...")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j: ", j)
		time.Sleep(time.Second)

	}

	nomes := [3]string{"teste1", "teste2", "teste3"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	usuario := map[string]string{
		"nome":      "paluan",
		"sobrenome": "batista",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}

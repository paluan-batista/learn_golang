package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	usuario := map[string]string{
		"nome":      "paluan",
		"sobrenome": "batista",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["sobrenome"])

}

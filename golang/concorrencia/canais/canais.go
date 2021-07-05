package main

import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}

func main() {
	canal := make(chan string)

	go escrever("testando canais", canal)
	for mensagem := range canal {

		fmt.Println(mensagem)
	}

	fmt.Println("FIM DA EXECUÇÂO")

}

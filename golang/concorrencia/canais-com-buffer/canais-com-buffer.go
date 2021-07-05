package main

import "fmt"

func main() {
	canal := make(chan string, 2) // canal com buffer

	canal <- "Testando canal com buffer"

	mensagem := <-canal
	fmt.Println(mensagem)
}

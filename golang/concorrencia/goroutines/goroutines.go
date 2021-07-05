package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("GOROUTINES")
	// CONCORRÊNCIA != PARALELISMO

	go escrever("Brincando com goroutine!") //goroutine
	escrever("Progamando em Go")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("WAIT GROUP")

	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		escrever("Brincando com goroutine - WAIT GROUP!")
		wait.Done()
	}()

	go func() {
		escrever("Progamando em Go")
		wait.Done()
	}()
	wait.Wait()

}

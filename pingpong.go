package main

import (
	"fmt"
	"time"
)

func jogador(nome string, entrada <-chan int, saida chan<- int) {
	for {

		_, ok := <-entrada
		if !ok {
			fmt.Println("Fim de jogo para", nome)
			return
		}

		fmt.Println(nome)
		time.Sleep(100 * time.Millisecond)

		saida <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go jogador("Ping", pong, ping)

	go jogador("Pong", ping, pong)

	ping <- 1

	close(ping)

}

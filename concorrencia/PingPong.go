package main

import (
	"fmt"
	"time"
)

func jogador(nome string, mesa_receber <-chan int, mesa_enviar chan<- int) {
	for {
		bola := <-mesa_receber
		fmt.Println(nome)
		time.Sleep(1 * time.Second)
		mesa_enviar <- bola
	}
}

func main() {
	pings := make(chan int)
	pongs := make(chan int)

	go jogador("Ping", pongs, pings)
	go jogador("Pong", pings, pongs)

	fmt.Println("O jogo vai iniciar...")
	pongs <- 1

	var entrada string
	fmt.Scanln(&entrada)

}

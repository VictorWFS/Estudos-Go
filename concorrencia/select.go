//funciona como um switch para canais
//select permite que você guarde operações de vários canais
// combinar goroutines e canais com select é um recurso poderoso do Go
// para o exemplo a seguir, serão dois canais.

package main

//cada canal receberá um valor após algum tempo, para simular, por exemplo,
//o bloqueio de operações RPC em execução em goroutines concorrentes

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "um"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)
		}
	}
}

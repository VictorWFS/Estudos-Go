//pacotes (caixas com funções)
//função contains
//contains: procura dentro de uma string outra string menor
//count o numero de vezes uma string menor aparece em uma string maior
//hasprefix verifica se o prefixo da palavra é o esperado retorna um boolean
//exemplo: terra - rr (true)

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("terra", "dor"))
	fmt.Println(strings.Count("Paralelepipedo", "e"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
}

//servidor client é o servidor relacionado ao que o usuário precisa
//o código abaixo vamos emitir uma solicitação de um cliente a um servidor http

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	//emitir uma requisição GET para um servidor
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Qual é o status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}

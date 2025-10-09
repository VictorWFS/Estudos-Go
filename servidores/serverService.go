//servidor service vai prestar um serviço ao usuário

package main

import (
	"fmt"
	"net/http"
)

// conceito principal em servidores net/http são as funções que estão no pacote http
func ola(w http.ResponseWriter, req *http.Request) {
	//as funções que servem como manipuladores recebem a http.ResponseWriter
	//http.Request recebe como argumentos, o gravador de resposta é usado para preencher
	//a resposta http
	fmt.Fprintf(w, "Ola\n")
}

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	//essa função lê todos os cabeçalhos de requisição HTTP
	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

func main() {
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalhos", cabecalhos)

	http.ListenAndServe(":8090", nil)
}

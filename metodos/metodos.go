package main

import "fmt"

/* type retangulo struct {
	comprimento, altura int
}

func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
} */

type pessoa struct {
	nome         string
	idade        int
	apresentacao string
}

func (a pessoa) apresentar() string {
	return a.apresentacao
}

func main() {
	/* r := retangulo{comprimento: 10, altura: 5}
	fmt.Println("Área: ", r.area())
	fmt.Println("Perimetro: ", r.perimetro()) */
	a := pessoa{nome: "Victor", idade: 21, apresentacao: "Olá me chamo Victor, tudo bem?"}
	fmt.Println(a.apresentar())
}

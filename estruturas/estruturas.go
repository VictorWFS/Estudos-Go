package main

import "fmt"

type carro struct {
	marca string
	nome  string
	valor float64
}

func main() {
	fmt.Println(carro{"Fiat", "Palio", 13.500})
}

//type () struct
//são coleçoes de campos
//agruapar dados
//formar registros

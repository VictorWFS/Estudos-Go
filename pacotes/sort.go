package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}

func (ps ParaNome) Less(i, j int) bool { //verifica se o item na posicao i é menor que o item na posição j
	return ps[i].Nome < ps[j].Nome
}

func (ps ParaNome) Swap(i, j int) { //troca os itens
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	criancas := []Dados{
		{"joao", 5},
		{"pedro", 10},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}

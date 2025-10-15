package main

import (
	"errors"
	"fmt"
)

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func Multiplicar(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func Subtrair(numeros ...int) int {
	if len(numeros) == 0 {
		return 0
	}

	total := numeros[0]
	for _, numero := range numeros[1:] {
		total = total - numero
	}
	return total
}

func Dividir(numeros ...float64) (float64, error) {
	total := numeros[0]

	for _, divisor := range numeros[1:] {
		if divisor == 0 {
			return 0, errors.New("erro: não é possível dividir por zero")
		}
		total = total / divisor
	}
	return total, nil
}

func main() {
	resultadoSoma := Soma(1, 5, 10)
	fmt.Println("Resultado da soma: ", resultadoSoma)

	resultadoMult := Multiplicar(5, 10)
	fmt.Println("Resultado da multiplicação: ", resultadoMult)

	resultadoSubrair := Subtrair(10, 5)
	fmt.Println("Resultado da subtração: ", resultadoSubrair)

	resultadoDivisão, err := Dividir(10, 4)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	} else {
		fmt.Println("Resultado da divisão: ", resultadoDivisão)
	}

}

package main

import "fmt"

const ebulicaoF = 212.0

func main() {
	tempF := ebulicaoF //operadores curtos :=, serve para declarar variável sem usar "var"
	tempC := (tempF - 32) * 5 / 9
	fmt.Printf("A temperatura de ebulição da água em ºF = %g,%T e a temperatura de ebulição da água em ºC =  %g,%T.", tempF, tempF, tempC, tempC)

}

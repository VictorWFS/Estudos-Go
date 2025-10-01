package main

import "fmt"

func main() {
	x := 9

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("Sim, x é multiplo de 2 ou 3")
	} else {
		fmt.Println("X não é multiplo de 2 nem 3.")
	}
}

// criar um código que verifique quais numeros entre 1 e 100
// são divisíveis por 3.
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Divisível por 3: ", i)
		}
	}
}

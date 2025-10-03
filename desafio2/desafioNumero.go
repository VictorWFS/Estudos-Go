//criar um código que entre numeros de 1 e 100, ao aparecer um multiplo
//de 3, deve aparecer "Pin" e multiplos de 5 "Pan"

package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Pin", i)
		} else if i%5 == 0 {
			fmt.Println("Pan", i)
		}
	}
}

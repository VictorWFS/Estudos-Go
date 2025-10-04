package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

//calcular a media de uma sala
func main() {
	lista := []float64{9.8, 9.3, 7.7, 8.2, 8.3}
	fmt.Println(media(lista))

	/* lista := []float64{9.8, 9.3, 7.7, 8.2, 8.3}
	total := 0.0
	for _, valor := range lista {
		total += valor
	}
	fmt.Println(total / float64(len(lista))) */
}

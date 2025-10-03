package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 5.3
	x[1] = 8
	x[2] = 4.2
	x[3] = 6.0
	x[4] = 7.8

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total = total + x[i]
	}

	fmt.Println(total / float64(len(x)))
}

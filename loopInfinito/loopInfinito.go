package main

import "fmt"

func main() {
	x := 0
	for {
		if x < 20 {
			x++
			fmt.Println("x < 20")
		} else {
			break
		}
	}

	for y := 0; y < 20; y++ {
		if y == 4 {
			continue
		}
		fmt.Println(y)

	}
}

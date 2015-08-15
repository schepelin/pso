package main

import (
	"fmt"
)

func parabola(x float64) float64 {
	return x * x
}

func main() {
	max := 10
	for i := 0; i < max; i++ {
		fmt.Print(i, "\n")
	}
}

package main

import (
	"fmt"
	"github.com/schepelin/pso/swarm"
)

func parabola(args ...float64) float64 {
	return args[0] * args[0]
}

func main() {
	extremum, coords := swarm.Run(
		parabola,
		1,
		5,
		-5.0,
		4.0,
	)
	fmt.Println(extremum, coords)
}

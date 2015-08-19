package main

import (
	"fmt"
	"github.com/schepelin/pso/swarm"
)

func parabola(args ...float64) float64 {
	return args[0] * args[0]
}

func main() {
	options := swarm.AlgorithmOptions{
		Func:            parabola,
		IterationsCount: 20,
		VariablesLength: 1,
		SwarmSize:       5,
		MinValue:        -5.0,
		MaxValue:        4.0,
	}
	extremum, coords := swarm.Run(options)
	fmt.Println(extremum, coords)
}

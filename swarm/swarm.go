package swarm

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const swarmPriority = 2.0
const individualPriority = 2.0

// AlgorithmOptions specifies initialization params for PSO
type AlgorithmOptions struct {
	Func            func(args ...float64) float64
	IterationsCount int
	VariablesLength int
	SwarmSize       int
	MinValue        float64
	MaxValue        float64
}

// Run start pso alghoritm for pased function with passed params
// returns founded extremum value and its coordinates
func Run(opts AlgorithmOptions) (float64, []float64) {

	rand.Seed(int64(time.Now().Nanosecond()))

	extremum := math.MaxFloat64
	swarmBest := make([]float64, opts.VariablesLength)
	swarm := make([]Particle, opts.SwarmSize)

	for i := 0; i < len(swarm); i++ {
		swarm[i].initialize(opts.VariablesLength, opts.MinValue, opts.MaxValue)
		particlePosition := swarm[i].position
		funcValue := opts.Func(particlePosition...)

		fmt.Println(particlePosition, funcValue)

		if funcValue < extremum {
			extremum = funcValue
			copy(swarmBest, particlePosition)
		}
	}
	fmt.Println(swarmBest)
	fmt.Println("-------------")
	for i := 0; i < opts.IterationsCount-1; i++ {

		for j := 0; j < len(swarm); j++ {
			swarm[j].update(&swarmBest, swarmPriority, individualPriority)
			particlePosition := swarm[j].position
			funcValue := opts.Func(particlePosition...)

			fmt.Println(particlePosition, funcValue)

			if funcValue < extremum {
				extremum = funcValue
				copy(swarmBest, particlePosition)
			}
		}
		fmt.Println(swarmBest)
		fmt.Println("-------------")
	}

	return extremum, swarmBest
}

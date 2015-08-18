package swarm

import (
	"math"
	"math/rand"
	"time"
)

// Run start pso alghoritm for pased function with passed params
func Run(f func(args ...float64) float64, variablesLength int,
	swarmSize int, minValue float64, maxValue float64) (float64, []float64) {
	rand.Seed(int64(time.Now().Nanosecond()))

	extremum := math.MaxFloat64
	swarmBest := make([]float64, variablesLength)
	swarm := make([]Particle, swarmSize)

	for i := 0; i < len(swarm); i++ {
		swarm[i].initialize(variablesLength, minValue, maxValue)
		particlePosition := swarm[i].position
		funcValue := f(particlePosition...)

		if funcValue < extremum {
			extremum = funcValue
			copy(swarmBest, particlePosition)
		}
	}

	return extremum, swarmBest
}

package pso

import (
	"math/rand"
)

// Particle represents basic partial in PSO alghoritm
type Particle struct {
	bestPosition []float64
	position     []float64
	velocity     []float64
}

type efficiencyFunction func([]float64) float64

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func (p *Particle) initialize(variablesLength int, minValue float64, maxValue float64) {

	p.position = make([]float64, variablesLength)
	p.bestPosition = make([]float64, variablesLength)
	p.velocity = make([]float64, variablesLength)

	for i := 0; i < variablesLength; i++ {
		value := random(minValue, maxValue)

		p.bestPosition[i] = value
		p.position[i] = value
	}

}

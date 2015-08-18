package swarm

import (
	"math/rand"
)

// Particle represents basic partial in PSO alghoritm
type Particle struct {
	bestPosition []float64
	position     []float64
	velocity     []float64
}

func random(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func (p *Particle) initialize(length int, minValue float64, maxValue float64) {
	p.position = make([]float64, length)
	p.bestPosition = make([]float64, length)
	p.velocity = make([]float64, length)

	for i := 0; i < length; i++ {
		value := random(minValue, maxValue)

		p.bestPosition[i] = value
		p.position[i] = value
	}
}

func (p *Particle) update(swarmBestPosition []float64, swarmPriority float64, individualPriority float64) {
	for i, value := range p.velocity {
		p.velocity[i] = value +
			swarmPriority*rand.Float64()*(p.bestPosition[i]-p.position[i]) +
			individualPriority*rand.Float64()*(swarmBestPosition[i]-p.position[i])
		p.position[i] = p.position[i] + p.velocity[i]
	}
}

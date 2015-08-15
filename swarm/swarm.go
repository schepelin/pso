package pso

import (
	"math/rand"
	"time"
)

// Run start pso alghoritm for pased function with passed params
func Run(f efficiencyFunction, variablesLength int, swarmSize int, minValue float64, maxValue float64) []float64 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return
}

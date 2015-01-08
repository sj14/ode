package ode

import "testing"
import "os"

var yStartSIR []float64 = []float64{}
var yStartPop []float64 = []float64{}

func TestMain(m *testing.M) {
	// SIR Start Values
	yStartSIR = append(yStartSIR, 700)
	yStartSIR = append(yStartSIR, 400)
	yStartSIR = append(yStartSIR, 100)

	// PopulationGrowthSimple Start Values
	yStartPop = append(yStartPop, 10)

	os.Exit(m.Run())
}

func BenchmarkEulerForwardPopulation(b *testing.B) {
	EulerForward(0, 0.1, 1000, yStartPop, populationGrowthSimple)
}

func BenchmarkRungeKuttaPopulation(b *testing.B) {
	RungeKutta4(0, 0.1, 1000, yStartPop, populationGrowthSimple)
}

func BenchmarkEulerForwardSIR(b *testing.B) {
	EulerForward(0, 0.1, 10000, yStartSIR, sir)
}

func BenchmarkRungeKuttaSIR(b *testing.B) {
	RungeKutta4(0, 0.1, 10000, yStartSIR, sir)
}

func populationGrowthSimple(t float64, y []float64) []float64 {
	result := make([]float64, 1)
	result[0] = 0.008 * y[0]
	return result
}

func sir(t float64, y []float64) []float64 {
	result := make([]float64, 3)
	result[0] = -0.0001 * y[1] * y[0]
	result[1] = 0.0001*y[1]*y[0] - 0.005*y[1]
	result[2] = 0.005 * y[1]
	return result
}

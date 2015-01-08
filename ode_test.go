// from fib_test.go
package ode

import "testing"
import "os"

var yStart []float64 = []float64{}

func TestMain(m *testing.M) {
	// SIR Start Values
	yStart = append(yStart, 700)
	yStart = append(yStart, 400)
	yStart = append(yStart, 100)
	
	os.Exit(m.Run())
}

func BenchmarkEulerForward(b *testing.B) {
	EulerForward(0, 0.1, 10000, yStart, sir)
}

func BenchmarkRungeKutta(b *testing.B) {
	RungeKutta4(0, 0.1, 10000, yStart, sir)
}

func sir(t float64, y []float64) []float64 {
	result := make([]float64, 3)
	result[0] = -0.0001 * y[1] * y[0]
	result[1] = 0.0001*y[1]*y[0] - 0.005*y[1]
	result[2] = 0.005 * y[1]
	return result
}

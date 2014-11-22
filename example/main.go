// Example to show how the ode package works
package main

import "fmt"
import "github.com/sj14/ode"

func main() {

	yStart := []float64{}

	// Population Growth Simple
	yStart = append(yStart, 10)

	/*
		// SIR
		yStart = append(yStart, 700)
		yStart = append(yStart, 400)
		yStart = append(yStart, 100)
	*/

	//y := ode.EulerForward(0, 1, 100, 0, yStart, populationGrowthSimple)
	y := ode.RungeKutta(0, 1, 100, 0, yStart, populationGrowthSimple)

	for _, val := range y {
		fmt.Println(val)
	}
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
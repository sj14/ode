// Example to show how the ode package works
package main

import "fmt"
import "github.com/sj14/ode"

func main() {

	yStartPop := []float64{}
	yStartSIR := []float64{}

	// Population Growth Simple Start Values
	yStartPop = append(yStartPop, 10)

	// SIR Start Values
	yStartSIR = append(yStartSIR, 700)
	yStartSIR = append(yStartSIR, 400)
	yStartSIR = append(yStartSIR, 100)

	eulerPop := ode.EulerForward(0, 10, 100, yStartPop, populationGrowthSimple)

	for _, val := range eulerPop {
		fmt.Println(val)
	}

	fmt.Println(yStartPop)

	rungePop := ode.RungeKutta4(0, 10, 100, yStartPop, populationGrowthSimple)

	for _, val := range rungePop {
		fmt.Println(val)
	}

	eulerSIR := ode.EulerForward(0, 10, 100, yStartSIR, sir)

	for _, val := range eulerSIR {
		fmt.Println(val)
	}

	rungeSIR := ode.RungeKutta4(0, 10, 100, yStartSIR, sir)

	for _, val := range rungeSIR {
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

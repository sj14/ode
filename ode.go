// Package ode provides the Euler-Forward and Runge-Kutta
// Method for solving ordinary differential equations.
package ode

import "fmt"

// BUG(sj14): The Runge-Kutta Method doesn't seem to work properly (not precise)

func EulerForward(from, h, to, t float64, y []float64, fn func(float64, []float64) []float64) [][]float64 {
	var steps int = int(to - from/h)
	var parameters = len(y)

	fmt.Println(steps)
	fmt.Println(parameters)

	// initialize 'outer slice'
	ySlice := make([][]float64, steps)
	// initialize first 'inner slice'
	ySlice[0] = make([]float64, parameters+1)

	// fill with start values
	for i := 0; i < parameters; i++ {
		ySlice[0][i+1] = y[i]
	}

	for step := 1; step < steps; step++ {
		t += h
		// initialize 'inner slice'
		ySlice[step] = make([]float64, parameters+1)
		ySlice[step][0] = t

		for value := 0; value < parameters; value++ {
			y[value] += h * fn(t, y)[value]
			ySlice[step][value+1] = y[value]
		}
	}
	return ySlice
}

func RungeKutta(from, h, to, t float64, y []float64, fn func(float64, []float64) []float64) [][]float64 {
	var steps int = int(to - from/h)
	var parameters = len(y)

	fmt.Println(steps)
	fmt.Println(parameters)

	// initialize 'outer slice'
	ySlice := make([][]float64, steps)
	// initialize first 'inner slice'
	ySlice[0] = make([]float64, parameters+1)

	// fill with start values
	for i := 0; i < parameters; i++ {
		ySlice[0][i+1] = y[i]
	}

	var k1 []float64 = make([]float64, parameters)
	var k2 []float64 = make([]float64, parameters)
	var k3 []float64 = make([]float64, parameters)
	var k4 []float64 = make([]float64, parameters)

	for step := 1; step < steps; step++ {
		// initialize 'inner slice'
		ySlice[step] = make([]float64, parameters+1)

		// TODO At the begin or end of the loop????
		for value := 0; value < parameters; value++ {
			y[value] += h * fn(t, y)[value]
		}

		for value := 0; value < parameters; value++ {
			k1[value] = h * 0.5 * fn(t, y)[value]
		}

		t += h / 2

		for value := 0; value < parameters; value++ {
			k2[value] = h * 0.5 * fn(t, k1)[value]
		}

		for value := 0; value < parameters; value++ {
			k3[value] = h * fn(t, k2)[value]
		}

		t += h / 2

		for value := 0; value < parameters; value++ {
			k4[value] = h * fn(t, k3)[value]
		}

		ySlice[step][0] = t

		for value := 0; value < parameters; value++ {
			ySlice[step][value+1] = y[value] + (k1[value]+2*k2[value]+2*k3[value]+k4[value])/6
		}
	}
	return ySlice
}

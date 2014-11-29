// Package ode provides the Euler-Forward and Runge-Kutta
// Method for solving ordinary differential equations.
package ode

import "fmt"

func EulerForward(from, h, to, t float64, y []float64, fn func(float64, []float64) []float64) [][]float64 {
	var steps int = int((to - from) / h)
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

	yn := make([]float64, parameters)

	for step := 1; step < steps; step++ {
		t += h
		// initialize 'inner slice'
		ySlice[step] = make([]float64, parameters+1)
		ySlice[step][0] = t

		for value := 0; value < parameters; value++ {
			yn[value] = ySlice[step-1][value+1]
		}

		for value := 0; value < parameters; value++ {
			y[value] += h * fn(t, yn)[value]
			ySlice[step][value+1] = y[value]
		}
	}
	return ySlice
}

// func RungeKutta4 is an implementation of the 4th order Runge-Kutta method
func RungeKutta4(from, h, to, t float64, y []float64, fn func(float64, []float64) []float64) [][]float64 {
	var steps int = int((to - from) / h)
	var parameters = len(y)

	fmt.Println(steps)
	fmt.Println(parameters)

	// initialize 'outer slice'
	ySlice := make([][]float64, steps+1)
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

	var yn []float64 = make([]float64, parameters)

	// the parameters to pass into the corresponding function calls
	var k2p []float64 = make([]float64, parameters)
	var k3p []float64 = make([]float64, parameters)
	var k4p []float64 = make([]float64, parameters)

	for step := 1; step <= steps; step++ {
		// initialize 'inner slice'
		ySlice[step] = make([]float64, parameters+1)

		// get last yn from the slice to make the rest of the function shorter
		for value := 0; value < parameters; value++ {
			yn[value] = ySlice[step-1][value+1]
		}

		// generate k1
		for value := 0; value < parameters; value++ {
			k1[value] = fn(t, yn)[value]
		}

		t += h / 2

		// generate the parameter for k2
		for value := 0; value < parameters; value++ {
			k2p[value] = yn[value] + 0.5*h*k1[value]
		}

		// generate k2
		for value := 0; value < parameters; value++ {
			k2[value] = fn(t, k2p)[value]
		}

		// generate the parameter for k3
		for value := 0; value < parameters; value++ {
			k3p[value] = yn[value] + 0.5*h*k2[value]
		}

		// generate k3
		for value := 0; value < parameters; value++ {
			k3[value] = fn(t, k3p)[value]
		}

		// generate the parameter for k4
		for value := 0; value < parameters; value++ {
			k4p[value] = yn[value] + h*fn(t, k3p)[value]
		}

		t += h / 2

		// generate k4
		for value := 0; value < parameters; value++ {
			k4[value] = fn(t, k4p)[value]
		}

		ySlice[step][0] = t

		// generate yn (saved in the slice)
		for value := 0; value < parameters; value++ {
			ySlice[step][value+1] = ySlice[step-1][value+1] + h*(k1[value]+2*k2[value]+2*k3[value]+k4[value])/6.0
		}
	}
	return ySlice
}

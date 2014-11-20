package main

import "fmt"

func main() {

	xStart := []float64{}
	fns := []func([]float64) float64{}

	/*
		// Population Growth Simple
		xStart = append(xStart, 10)
		fns = append(fns, populationGrowthSimple)
	*/

	// SIR
	xStart = append(xStart, 700)
	xStart = append(xStart, 400)
	xStart = append(xStart, 100)

	fns = append(fns, dz1)
	fns = append(fns, dz2)
	fns = append(fns, dz3)

	//x := EulerForward(0, 1, 100, 0, xStart, fns)
	x := RungeKutta(0, 1, 100, 0, xStart, fns)

	for _, val := range x {
		fmt.Println(val)
	}
}

func populationGrowthSimple(x []float64) float64 {
	var alpha float64 = 0.008
	return alpha * x[0]
}

func dz1(y []float64) float64 {
	return -0.0001 * y[1] * y[0]
}

func dz2(y []float64) float64 {
	return 0.0001*y[1]*y[0] - 0.005*y[1]
}

func dz3(y []float64) float64 {
	return 0.005 * y[1]
}

// EulerForward this should produce a comment
func EulerForward(from, h, to, t float64, x []float64, fn []func([]float64) float64) [][]float64 {

	var steps int = int(to - from/h)
	var parameters = len(x)

	fmt.Println(steps)
	fmt.Println(parameters)

	xSlice := make([][]float64, steps)        // initialize 'outer slice'
	xSlice[0] = make([]float64, parameters+1) // initialize first 'inner slice'

	// fill with start values
	for i := 1; i <= parameters; i++ {
		xSlice[0][i] = x[i-1]
	}

	for step := 1; step < steps; step++ {
		t += h
		//fmt.Printf("t: %v\n", t)
		xSlice[step] = make([]float64, parameters+1) // initialize 'inner slice'
		xSlice[step][0] = t

		for value := 0; value < parameters; value++ {
			x[value] += h * fn[value](x)
			xSlice[step][value+1] = x[value]
			//fmt.Printf("x: %v\n", x[value])
		}
	}
	return xSlice
}

func RungeKutta(from, h, to, t float64, x []float64, fn []func([]float64) float64) [][]float64 {

	var steps int = int(to - from/h)
	var parameters = len(x)

	fmt.Println(steps)
	fmt.Println(parameters)

	xSlice := make([][]float64, steps)        // initialize 'outer slice'
	xSlice[0] = make([]float64, parameters+1) // initialize first 'inner slice'

	// fill with start values
	for i := 1; i <= parameters; i++ {
		xSlice[0][i] = x[i-1]
	}

	for step := 1; step < steps; step++ {
		t += h
		//fmt.Printf("t: %v\n", t)
		xSlice[step] = make([]float64, parameters+1) // initialize 'inner slice'
		xSlice[step][0] = t

		for value := 0; value < parameters; value++ {
			
			x[value] += h * fn[value](x)
			
			var k1 []float64 = make([]float64, parameters)

			k1[0] = x[value] + h*0.5*fn[value](x)
			k1[1] = x[value] + h*0.5*fn[value](x)
			k1[2] = x[value] + h*0.5*fn[value](x)

			var k2 []float64 = make([]float64, parameters)

			k2[0] = x[value] + h*0.5*fn[value](k1)
			k2[1] = x[value] + h*0.5*fn[value](k1)
			k2[2] = x[value] + h*0.5*fn[value](k1)

			var k3 []float64 = make([]float64, parameters)

			k3[0] = x[value] + h*fn[value](k2)
			k3[1] = x[value] + h*fn[value](k2)
			k3[2] = x[value] + h*fn[value](k2)

			var k4 []float64 = make([]float64, parameters)

			k4[0] = x[value] + h*fn[value](k3)
			k4[1] = x[value] + h*fn[value](k3)
			k4[2] = x[value] + h*fn[value](k3)
			
			

			xSlice[step][value+1] = h*(k1[0] + 2*k2[0] + 2*k3[0] + k4[0]) / 6
			
			

		}
	}
	return xSlice
}

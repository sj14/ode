package main

import "fmt"

func main() {

	/*
	xStart := []float64{}
	xStart = append(xStart, 10)
	fns := []func([]float64) float64{}
	fns = append(fns, populationGrowthSimple)
	EulerForward(0, 1, 100, 0, xStart, fns)
	*/
	
	xStart := []float64{}
	xStart = append(xStart, 700)
	xStart = append(xStart, 400)
	xStart = append(xStart, 100)
	
	fns := []func([]float64) float64{}
	fns = append(fns, dz1)
	fns = append(fns, dz2)
	fns = append(fns, dz3)
	
	_, x := EulerForward(0, 1, 100, 0, xStart, fns)
	
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
	return 0.0001*y[1]*y[0] - 0.005*y[2]
}

func dz3(y []float64) float64 {
	return 0.005 * y[1]
}

func EulerForward(from, h, to, t float64, x []float64, fn []func([]float64) float64) ([]float64, [][]float64) {

	var steps int = int(to - from/h)
	var parameters = len(x)

	fmt.Println(steps)
	fmt.Println(parameters)

	tSlice := make([]float64, steps)
	xSlice := make([][]float64, steps) // initialize 'outer slice'

	for step := 0; step < steps; step++ {
		t += h
		fmt.Printf("t: %v\n", t)
		xSlice[step] = make([]float64, parameters) 	// initialize 'inner slice'

		for value := 0; value < parameters; value++ {
			x[value] += h * fn[value](x)
			xSlice[step][value] = x[value]
			fmt.Printf("x: %v\n", x[value])
			
		}
		tSlice = append(tSlice, t)
	}

	return tSlice, xSlice
}

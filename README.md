ode
===
[![GoDoc](https://godoc.org/github.com/sj14/ode?status.png)](https://godoc.org/github.com/sj14/ode) [![Build Status](https://travis-ci.org/sj14/ode.svg)](https://travis-ci.org/sj14/ode)

A package for the go programming language to solve ordinary differential equations.

## Requirements

- Go >= 1.4

## Example
```go
// Example to show how the ode package works
package main

import "fmt"
import "github.com/sj14/ode"

func main() {
  // SIR Start Values
  yStartSIR := []float64{700, 400, 100}

  // Do the calculation
  y := ode.RungeKutta4(0, 10, 100, yStartSIR, sir)

  // Output the results to the console
  for _, val := range y {
    fmt.Println(val)
  }
}

func sir(t float64, y []float64) []float64 {
  result := make([]float64, 3)
  result[0] = -0.0001 * y[1] * y[0]
  result[1] = 0.0001*y[1]*y[0] - 0.005*y[1]
  result[2] = 0.005 * y[1]
  return result
}
```
## Output
```go
[0 700 400 100]
[10 410.90567125203955 662.4382350812937 126.65609366666666]
[20 191.97505869594923 843.2048493089754 164.82009199507524]
[30 79.86753851144415 911.0287568497188 209.10370463883703]
[40 32.32376388920436 912.7706547315336 254.90558137926206]
[50 13.263171694988241 886.7634313299168 299.973396975095]
[60 5.607931908716353 850.9501619076912 343.4419061835925]
[70 2.4571403113927808 812.5094770189563 385.03338266965085]
[80 1.1171479813071248 774.1848647682339 424.697987250459]
[90 0.5268422655752388 737.0010917960352 462.4720659383895]
[100 0.2574396186972187 701.3189883896745 498.4235719916282]
```

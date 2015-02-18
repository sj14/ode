package ode_test

import (
	"fmt"
	"testing"

	"github.com/sj14/ode"
)

var yStartSIR = []float64{}
var yStartPop = []float64{}

func init() {
	// SIR Start Values
	yStartSIR = append(yStartSIR, 700)
	yStartSIR = append(yStartSIR, 400)
	yStartSIR = append(yStartSIR, 100)

	// PopulationGrowthSimple Start Values
	yStartPop = append(yStartPop, 10)
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

func BenchmarkEulerForwardPopulation(b *testing.B) {
	ode.EulerForward(0, 0.1, 10000, yStartPop, populationGrowthSimple)
}

func BenchmarkRungeKuttaPopulation(b *testing.B) {
	ode.RungeKutta4(0, 0.1, 10000, yStartPop, populationGrowthSimple)
}

func BenchmarkEulerForwardSIR(b *testing.B) {
	ode.EulerForward(0, 0.1, 10000, yStartSIR, sir)
}

func BenchmarkRungeKuttaSIR(b *testing.B) {
	ode.RungeKutta4(0, 0.1, 10000, yStartSIR, sir)
}

func ExampleEulerForward_population() {
	y := ode.EulerForward(0, 10, 100, yStartPop, populationGrowthSimple)

	for _, val := range y {
		fmt.Println(val)
	}
	// Output:
	// [0 10]
	// [10 10.8]
	// [20 11.664000000000001]
	// [30 12.597120000000002]
	// [40 13.604889600000002]
	// [50 14.693280768000001]
	// [60 15.868743229440001]
	// [70 17.138242687795202]
	// [80 18.50930210281882]
	// [90 19.990046271044324]
	// [100 21.58924997272787]
}

func ExampleEulerForward_sir() {
	y := ode.EulerForward(0, 10, 100, yStartSIR, sir)

	for _, val := range y {
		fmt.Println(val)
	}
	// Output:
	// [0 700 400 100]
	// [10 420 660 120]
	// [20 142.79999999999995 904.2 153]
	// [30 13.68023999999997 988.10976 198.21]
	// [40 0.16266133685759776 952.2218506631424 247.61548800000003]
	// [50 0.00777165764371518 904.7656478091992 295.2265805331572]
	// [60 0.0007401287811479003 859.5343969476019 340.46486292361715]
	// [70 0.00010396263558037628 816.5583132663673 383.44158277099723]
	// [80 1.9071081228138196e-05 775.7304824946034 424.2694984343156]
	// [90 4.277062185340778e-06 736.9439731638922 463.0560225590458]
	// [100 1.1251069850067054e-06 700.0967776576529 499.9032212172404]
}

func ExampleRungeKutta4_population() {
	yz := ode.RungeKutta4(0, 10, 100, yStartPop, populationGrowthSimple)

	for _, value := range yz {
		fmt.Println(value)
	}
	// Output:
	// [0 10]
	// [10 10.832870400000001]
	// [20 11.735108110319617]
	// [30 12.71249052890813]
	// [40 13.771276236088923]
	// [50 14.91824507081511]
	// [60 16.16074154475789]
	// [70 17.506721872225803]
	// [80 18.96480491706675]
	// [90 20.544327382786687]
	// [100 22.255403599289938]
}

func ExampleRungeKutta4_sir() {
	y := ode.RungeKutta4(0, 10, 100, yStartSIR, sir)

	for _, val := range y {
		fmt.Println(val)
	}
	// Output:
	// [0 700 400 100]
	// [10 410.90567125203955 662.4382350812937 126.65609366666666]
	// [20 191.97505869594923 843.2048493089754 164.82009199507524]
	// [30 79.86753851144415 911.0287568497188 209.10370463883703]
	// [40 32.32376388920436 912.7706547315336 254.90558137926206]
	// [50 13.263171694988241 886.7634313299168 299.973396975095]
	// [60 5.607931908716353 850.9501619076912 343.4419061835925]
	// [70 2.4571403113927808 812.5094770189563 385.03338266965085]
	// [80 1.1171479813071248 774.1848647682339 424.697987250459]
	// [90 0.5268422655752388 737.0010917960352 462.4720659383895]
	// [100 0.2574396186972187 701.3189883896745 498.4235719916282]
}

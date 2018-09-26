package main

import (
	"fmt"
)

const precision float64 = 0.0001

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	var z1 float64 = 1.7
	for {
		z1 -= (z*z - x) / (2 * z)
		if (z1-z)*(z1-z) < precision {
			return z1
		}
		z = z1
	}
	return 0
}

func main() {
	fmt.Println(Sqrt(2))
}

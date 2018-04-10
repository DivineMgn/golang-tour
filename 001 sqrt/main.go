// https://go-tour-ru-ru.appspot.com/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

const (
	eps     = 1e-15
	maxIter = 100
)

func main() {
	fmt.Println(Sqrt(2))
}

// Sqrt - function returns the square root of a number
func Sqrt(x float64) float64 {
	z1 := float64(0.1)

	for i := 0; i < maxIter; i++ {
		z2 := z1 - (z1*z1-x)/(2*z1)

		if math.Abs(z1-z2) < eps {
			break
		}

		z1 = z2
	}

	return z1
}

// https://go-tour-ru-ru.appspot.com/methods/20

package main

import (
	"fmt"
	"math"
)

const (
	eps     = 1e-15
	maxIter = 100
)

// ErrNegativeSqrt - error type
type ErrNegativeSqrt float64

func main() {
	if x, err := Sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}

	if x, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}

func (ex ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(ex))
}

// Sqrt - function returns the square root of a number or error
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z1 := float64(0.1)

	for i := 0; i < maxIter; i++ {
		z2 := z1 - (z1*z1-x)/(2*z1)

		if math.Abs(z1-z2) < eps {
			break
		}

		z1 = z2
	}

	return z1, nil
}

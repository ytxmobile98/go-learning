package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	const iterations = 100

	z := 1.0
	for i := 0; i < iterations; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func ErrNegativeSqrt(x float64) error {
	return fmt.Errorf("cannot Sqrt negative number: %v", x)
}

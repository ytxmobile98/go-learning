package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	const iterations = 100

	z := 1.0
	for i := 0; i < iterations; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

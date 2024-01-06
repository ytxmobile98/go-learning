package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	a, b := 0, 1

	return func() int {
		var result int

		if (i == 0) || (i == 1) {
			result = i
		} else {
			result = a + b
			a = b
			b = result
		}
		i++

		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

package main

import (
	"fmt"
)

type Result struct {
	Odd  int
	Even int
}

func IsEven(n int) bool {
	return n%2 == 0
}

func CountOddEven(low, high int) (result Result) {
	const maxConcurrency = 100

	channels := struct {
		odd  chan int
		even chan int
	}{
		odd:  make(chan int, maxConcurrency),
		even: make(chan int, maxConcurrency),
	}

	go func() {
		concurrency := 0

		for i := low; i <= high; {
			if concurrency >= maxConcurrency {
				continue
			}

			concurrency++
			go func(i int) {
				if IsEven(i) {
					channels.even <- i
				} else {
					channels.odd <- i
				}

				concurrency--
			}(i)
			i++
		}
	}()

	for total := high - low + 1; result.Even+result.Odd < total; {
		select {
		case <-channels.even:
			result.Even++
		case <-channels.odd:
			result.Odd++
		}
	}
	return
}

func main() {
	result := CountOddEven(10, 100000)

	fmt.Println("Odd:", result.Odd)
	fmt.Println("Even:", result.Even)
}

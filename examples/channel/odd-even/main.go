package main

import (
	"fmt"
	"sync/atomic"
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
		concurrency := int64(0)

		for i := low; i <= high; {
			if concurrency >= maxConcurrency {
				continue
			} else {
				atomic.AddInt64(&concurrency, 1)
				go func(i int) {
					if IsEven(i) {
						channels.even <- i
					} else {
						channels.odd <- i
					}

					atomic.AddInt64(&concurrency, -1)
				}(i)
				i++
			}
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
	result := CountOddEven(10, 1000)

	fmt.Println("Odd:", result.Odd)
	fmt.Println("Even:", result.Even)
}

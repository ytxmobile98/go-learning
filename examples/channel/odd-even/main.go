package main

type Result struct {
	Odd  int
	Even int
}

func IsEven(n int) bool {
	return n%2 == 0
}

func CountOddEven(low, high int) (result Result) {
	channels := struct {
		odd  chan int
		even chan int
	}{
		odd:  make(chan int, 10),
		even: make(chan int, 10),
	}

	for i := low; i <= high; i++ {
		go func() {
			if IsEven(i) {
				channels.even <- i
			} else {
				channels.odd <- i
			}
		}()
	}

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
	result := CountOddEven(10, 100)

	println("Odd:", result.Odd)
	println("Even:", result.Even)
}

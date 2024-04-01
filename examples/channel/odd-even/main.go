package main

type Result struct {
	Odd  int
	Even int
}

func IsEven(n int) bool {
	return n%2 == 0
}

func CountOddEven(low, high int) (result Result) {
	for i := low; i <= high; i++ {
		if IsEven(i) {
			result.Even++
		} else {
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

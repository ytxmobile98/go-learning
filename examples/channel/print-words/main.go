package main

import "fmt"

func main() {
	maxCount := 100
	words := []string{"cat", "fish", "dog"}

	for i := 0; i < maxCount; i++ {
		for _, word := range words {
			fmt.Println(word)
		}
	}
}

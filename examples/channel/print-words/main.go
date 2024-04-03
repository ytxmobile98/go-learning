package main

import (
	"fmt"
)

func InitChannels[T any](count int) []chan T {
	channels := make([]chan T, count)
	for i := range channels {
		channels[i] = make(chan T)
	}
	return channels
}

func main() {
	maxCount := 100

	words := []string{"cat", "fish", "dog"}
	channels := InitChannels[bool](len(words))

	done := 0

	print := func(wordIdx int) {
		word := words[wordIdx]
		send := (chan<- bool)(channels[(wordIdx+1)%len(words)])
		recv := (<-chan bool)(channels[wordIdx])

		for i := 0; i < maxCount; i++ {
			<-recv
			fmt.Println(word)

			if i < maxCount-1 {
				send <- true
			} else {
				close(send)
				done++
			}
		}
	}

	for i := range words {
		go print(i)
	}
	channels[0] <- true // initialize the first print

	for done < len(words) {
		// wait for all prints to finish
	}
}

package main

import (
	"fmt"
	"sync"
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

	wg := sync.WaitGroup{}
	wg.Add(len(words))

	print := func(word string, recv <-chan bool, send chan<- bool) {
		for i := 0; i < maxCount; i++ {
			<-recv
			fmt.Println(word)
			if i == maxCount-1 {
				close(send)
				wg.Done()
			} else {
				send <- true
			}
		}
	}

	for i, word := range words {
		go print(word, channels[i], channels[(i+1)%len(words)])
	}
	channels[0] <- true // initialize the first print

	wg.Wait()
}

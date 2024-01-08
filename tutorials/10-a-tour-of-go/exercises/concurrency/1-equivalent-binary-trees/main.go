package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)

	ch2 := make(chan int)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 || v1 != v2 {
			return false
		} else if !ok1 && !ok2 {
			return true
		}
	}
}

func main() {
	pairs := [][2]int{
		{1, 1},
		{1, 2},
		{2, 2},
		{2, 3},
	}

	for _, pair := range pairs {
		t1 := tree.New(pair[0])
		t2 := tree.New(pair[1])

		fmt.Printf("Same(tree.New(%d), tree.New(%d)): %v\n",
			pair[0], pair[1], Same(t1, t2))
	}
}

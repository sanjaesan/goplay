package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	fmt.Println(".....Equivalent Binary Trees Solution.....")
	treeOne := tree.New(5)
	treeTwo := tree.New(5)
	fmt.Println("Are the trees the same?\n", Same(treeOne, treeTwo))
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// exits func as early as possible
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var result bool
	done := make(chan bool)
	channelA := make(chan int)
	channelB := make(chan int)
	go Walk(t1, channelA)
	go Walk(t2, channelB)
	go func() {
		for i := range channelA {
			if i == <-channelB {
				result = true
			} else {
				result = false
				break
			}
		}
		done <- true
	}()
	<-done
	return result
}

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var out = make(chan int)
	var in = make(chan int)

	go mul(in, out)
	in <- 2
	output := <-out
	fmt.Println(output)

	channelA := createChan(1, 2, 3, 4, 5, 6, 7, 8, 9)
	channelB := createChan(11, 12, 13, 14, 15, 16, 17, 18, 19)
	channelC := createChan(21, 22, 23, 24, 25, 26, 27, 28, 29)

	for val := range merge(channelA, channelB, channelC) {
		fmt.Println(val)
	}
}

// createChan returns a channel that signals int type ...
// used to create a new channel for testing the merge function
func createChan(numbers ...int) <-chan int {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(numbers))
	for _, number := range numbers {
		go func(number int) {
			ch <- number
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}(number)
		wg.Done()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

// illustrates directions on channels
func mul(in <-chan int, out chan<- int) {
	fmt.Println("Initializing go routine ...")
	num := <-in
	result := num * 2
	out <- result
}

// Takes varying number of channels and merge
func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(channel <-chan int) {
			for v := range channel {
				out <- v
			}
			wg.Done()
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

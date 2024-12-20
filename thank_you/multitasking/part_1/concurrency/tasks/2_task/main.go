package main

import (
	"fmt"
	"sync"
	"time"
)

func rangeGen(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			time.Sleep(50 * time.Millisecond)
			out <- i
		}
	}()
	return out
}

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		ch := ch
		go func() {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in1 := rangeGen(11, 15)
	in2 := rangeGen(21, 25)
	in3 := rangeGen(31, 35)

	start := time.Now()
	merged := merge(in1, in2, in3)
	for val := range merged {
		fmt.Print(val, " ")
	}
	fmt.Println()
	fmt.Println("Took", time.Since(start))
}

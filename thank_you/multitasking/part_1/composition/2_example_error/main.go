package main

import (
	"fmt"
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

func merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in1 {
			out <- val
		}
		for val := range in2 {
			out <- val
		}
	}()
	return out
}

func main() {
	start := time.Now()

	in1 := rangeGen(11, 15)
	in2 := rangeGen(21, 25)

	merged := merge(in1, in2)
	for val := range merged {
		fmt.Print(val, " ")
	}
	fmt.Println()
	fmt.Println("Took", time.Since(start))
}

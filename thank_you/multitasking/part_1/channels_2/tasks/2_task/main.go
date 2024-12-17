package main

import (
	"fmt"
)

func work(done chan struct{}, out chan int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	done <- struct{}{}
}

func main() {
	out := make(chan int)
	done := make(chan struct{})

	go work(done, out) // (1)

	go func() {
		<-done // (2)
		fmt.Println("Done! closing channel...")
		defer close(out)
	}()

	for n := range out { // (3)
		fmt.Println(n)
	}
}

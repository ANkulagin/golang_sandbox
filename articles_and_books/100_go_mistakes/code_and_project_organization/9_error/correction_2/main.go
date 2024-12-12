package main

import "fmt"

func merge[T any](ch1, ch2 <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for {
			select {
			case val, ok := <-ch1:
				if ok {
					out <- val
				} else {
					ch1 = nil
				}
			case val, ok := <-ch2:
				if ok {
					out <- val
				} else {
					ch2 = nil
				}
			}
			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()
	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	merged := merge(ch1, ch2)
	for val := range merged {
		fmt.Printf("Value: %\n", val)
	}
}

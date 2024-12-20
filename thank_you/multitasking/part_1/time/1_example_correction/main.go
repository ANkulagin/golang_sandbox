package main

import "fmt"

func rangeGen(cancel <-chan struct{}, start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		for i := start; i < stop; i++ {
			select {
			case <-cancel:
				return
			case out <- i:
			}
		}
		close(out)
	}()
	return out
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	generated := rangeGen(cancel, 1, 10)
	for n := range generated {
		if n == 2 {
			break
		}
		fmt.Println(n)
	}
}

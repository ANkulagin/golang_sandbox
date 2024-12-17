package main

import "fmt"

func main() {
	var stream chan int
	fmt.Printf("%v\n", stream)
	go func() {
		stream <- 1
	}()

	//	<-stream
	// fatal error: all goroutines are asleep - deadlock!
}

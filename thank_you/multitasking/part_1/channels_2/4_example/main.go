package main

import "fmt"

func main() {
	stream := make(chan int, 2)
	stream <- 1
	stream <- 2
	close(stream)

	val, ok := <-stream
	fmt.Println(val, ok)
	// 1 true

	val, ok = <-stream
	fmt.Println(val, ok)
	// 2 true

	val, ok = <-stream
	fmt.Println(val, ok)
	// 0 false
}

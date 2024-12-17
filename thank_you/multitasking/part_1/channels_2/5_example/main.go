package main

import (
	"fmt"
	"time"
)

func main() {
	stream := make(chan int, 3)
	fmt.Printf("%v\n", stream)
	go func() {
		fmt.Println("Sending...")
		stream <- 1
		stream <- 2
		stream <- 3
		close(stream)
		fmt.Println("Sent and closed!")
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Receiving...")
	for val := range stream {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
	fmt.Println("Received!")
}

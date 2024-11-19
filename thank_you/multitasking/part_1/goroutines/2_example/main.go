package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		fmt.Println("B: Sending message...") // (2)
		messages <- "ping"                   // (3)
		fmt.Println("B: Message sent!")      // (6)
	}()

	fmt.Println("A: Doing some work...") // (1)
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("A: Ready to receive a message...") // (4)

	<-messages //  (5)

	fmt.Println("A: Messege received!") // (5)
	time.Sleep(10000 * time.Millisecond)
}

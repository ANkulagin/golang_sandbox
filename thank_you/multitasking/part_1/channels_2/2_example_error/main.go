package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	stream := make(chan bool)

	send := func() {
		defer wg.Done()
		fmt.Println("Sender ready to send...")
		stream <- true // (1)
		fmt.Println("Sent!")
	}

	receive := func() {
		defer wg.Done()
		fmt.Println("Receiver not ready yet...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Receiver ready to receive...")
		<-stream // (2)
		fmt.Println("Received!")
	}

	go send()
	go receive()
	wg.Wait()
}

// при запуске будет видно, что ожидаемый вывод не логичен

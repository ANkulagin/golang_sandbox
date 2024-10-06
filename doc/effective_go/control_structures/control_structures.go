package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	number := rand.Int() % 10
	// If statement
	if x := number; x > 5 {
		fmt.Println("x is greater than 5")
	}

	// For loop (C-like)
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// For loop (range)
	arr := []int{1, 2, 3, 4}
	for k, val := range arr {
		fmt.Println("Index:", k)
		fmt.Println("Value:", val)
	}

	// Infinite loop with break
	i := 0
	for {
		fmt.Println("Looping...", i)
		i++
		if i == 3 {
			break
		}
	}

	// Switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the week")
	default:
		fmt.Println("Midweek")
	}

	// Type switch
	var t interface{}
	t = 42
	switch t := t.(type) {
	case int:
		fmt.Printf("Type is int: %d\n", t)
	case string:
		fmt.Printf("Type is string: %s\n", t)
	default:
		fmt.Printf("Unknown type: %T\n", t)
	}

	// Select
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "Hello from chan1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Hello from chan2"
	}()

	time.Sleep(1 * time.Second)
	select {
	case msg := <-chan1:
		fmt.Println("Received:", msg)
	case msg := <-chan2:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received")
	}
}

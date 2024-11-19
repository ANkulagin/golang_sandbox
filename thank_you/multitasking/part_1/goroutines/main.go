package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		say(1, "good day")
		defer wg.Done()
	}()

	go func() {
		say(2, "bad food")
		defer wg.Done()
	}()

	wg.Wait()
}

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

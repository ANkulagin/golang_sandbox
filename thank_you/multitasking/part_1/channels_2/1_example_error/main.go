package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func say(id int, phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func main() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}
	for idx, phrase := range phrases {
		go say(idx+1, phrase)
	}
}

// программа ничего не успеет сделать

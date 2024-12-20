package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "one,two,,four"

	in := make(chan string)
	go func() { // (1)
		words := strings.Split(str, ",")
		for _, word := range words {
			in <- word
		}
		close(in)
	}()

	for word := range in {
		if word != "" {
			fmt.Printf("%s ", word)
		}
	}
}

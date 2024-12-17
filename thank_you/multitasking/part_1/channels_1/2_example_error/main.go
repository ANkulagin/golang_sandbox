package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "one,two,,four"
	stream := make(chan string)
	go submit(str, stream)
	printStr(stream)
}

func submit(str string, stream chan string) {
	words := strings.Split(str, ",")
	for _, word := range words {
		stream <- word
	}
	<-stream // (1)допустим мы ошиблись
	close(stream)
}

func printStr(stream chan string) {
	for word := range stream {
		if word != "" {
			fmt.Printf("%s ", word)
		}
	}
	fmt.Println()
}

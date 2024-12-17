package main

import (
	"fmt"
)

func encode(src string) string {
	input := make(chan string)
	encoded := make(chan string)
	output := make(chan string)

	go submitter(src, input)
	go encoder(input, encoded)
	go receiver(encoded, output)

	return <-output
}

func submitter(src string, input chan<- string) {
	input <- src
	close(input)
}

func encoder(input <-chan string, encoded chan<- string) {
	encoded <- encodeWord(<-input)
	close(encoded)
	// Реализуйте логику шифрования и передачи в канал encoded
}

func receiver(encoded <-chan string, output chan<- string) {

	output <- <-encoded
	close(output)
	// Реализуйте логику передачи результата в канал output
}

func encodeWord(word string) string {
	wordEncrypted := word + "!"
	return wordEncrypted
}

func main() {
	src := "go is awesome"
	res := encode(src)
	fmt.Println(res)
}

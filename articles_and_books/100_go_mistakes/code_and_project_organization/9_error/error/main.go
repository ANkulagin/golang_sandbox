package main

import "io"

// Функция, которая принимает io. Writer с использованием дженериков
func foo[T io.Writer](w T) {
	b := []byte("Hello, Go!")
	_, _ = w.Write(b)
}

func main() {
	var writer io.Writer = nil
	foo(writer)
}

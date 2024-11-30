package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

type counter map[string]int

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	stats := make(counter)

	// начало решения
	messages := make(chan struct {
		word  string
		count int
	})

	var wg sync.WaitGroup
	wg.Add(1)
	// Запустите одну горутину.
	// В горутине пройдите по словам, посчитайте количество цифр в каждом,
	// и запишите результаты в канал counted.
	// В основной функции считайте значения из канала и заполните stats.
	go func() {
		for _, word := range words {
			count := countDigits(word)
			messages <- struct {
				word  string
				count int
			}{word, count}
		}
		defer wg.Done()
		close(messages)
	}()

	for message := range messages {
		stats[message.word] = message.count
	}
	wg.Wait()

	// конец решения
	return stats

}

func main() {
	input := "Hello 123 world 4567"
	stats := countDigitsInWords(input)
	for word, count := range stats {
		fmt.Printf("Word: %s, Digits: %d\n", word, count)
	}
}

/*
Word: Hello, Digits: 0
Word: 123, Digits: 3
Word: world, Digits: 0
Word: 4567, Digits: 4
*/

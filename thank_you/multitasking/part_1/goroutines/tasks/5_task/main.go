package main

import (
	"fmt"
	"unicode"
)

type pair struct {
	word  string
	count int
}

type counter map[string]int

type nextFunc func() string

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func submitWords(next nextFunc, pending chan string) {
	// Реализуйте логику передачи слов в канал pending
	for {
		word := next()
		if word == "" {
			break
		}
		pending <- word
	}
	close(pending)
}

func countWords(pending chan string, counted chan pair) {
	// Реализуйте логику подсчета цифр и передачи результатов в канал counted
	for word := range pending {
		count := countDigits(word)
		counted <- pair{word, count}
	}
	close(counted)
}

func fillStats(counted chan pair) counter {
	stats := make(counter)
	// Реализуйте логику заполнения stats из канала counted
	for result := range counted {
		stats[result.word] = result.count
	}
	return stats
}

func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)
}

func main() {
	words := []string{"Hello", "123", "world", "4567"}
	index := 0

	// Функция-генератор для следующего слова
	next := func() string {
		if index >= len(words) {
			return "" // Возвращаем пустую строку, чтобы обозначить окончание генерации
		}
		word := words[index]
		index++
		return word
	}

	stats := countDigitsInWords(next)
	for word, count := range stats {
		fmt.Printf("Word: %s, Digits: %d\n", word, count)
	}
}

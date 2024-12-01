/*
Когда канал pending закрывается, цикл for (56) завершается, и канал counted закрывается. Цикл for (63) продолжает работать, пока канал counted не будет закрыт
*/
package main

import (
	"fmt"
	"sync"
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

func countDigitsInWords(next nextFunc) counter {
	stats := make(counter)
	pending := make(chan string)
	counted := make(chan pair)

	var wg sync.WaitGroup

	// Читатель: отправляет слова в канал pending
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			word := next()
			if word == "" {
				break
			}
			pending <- word
		}
		close(pending)
	}()

	// Счетовод: считывает слова из pending, считает цифры и отправляет в counted
	wg.Add(1)
	go func() {
		defer wg.Done()
		for word := range pending {
			count := countDigits(word)
			counted <- pair{word, count}
		}
		close(counted)
	}()

	for result := range counted {
		stats[result.word] = result.count
	}

	return stats
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

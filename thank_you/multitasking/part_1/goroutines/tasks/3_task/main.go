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
	counted := make(chan pair)

	var wg sync.WaitGroup
	wg.Add(1)

	// начало решения
	go func() {
		defer wg.Done()
		for {
			word := next()
			if word == "" { // Окончание генератора
				break
			}
			count := countDigits(word)
			counted <- pair{word, count}
		}
		close(counted)
	}()

	for result := range counted {
		stats[result.word] = result.count
	}
	// Запустите одну горутину.
	// В горутине пройдите по словам, посчитайте количество цифр в каждом,
	// и запишите результаты в канал counted.

	// В основной функции считайте значения из канала и заполните stats.

	// конец решения

	wg.Wait()
	return stats
}

func main() {
	words := []string{"Hello", "123", "world", "4567"}
	count := 0

	// Функция-генератор для следующего слова
	next := func() string {
		if count >= len(words) {
			return "" // Возвращаем пустую строку, чтобы обозначить окончание генерации
		}
		word := words[count]
		count++
		return word
	}

	stats := countDigitsInWords(next)
	for word, count := range stats {
		fmt.Printf("Word: %s, Digits: %d\n", word, count)
	}
}

/*
Word: 4567, Digits: 4
Word: Hello, Digits: 0
Word: 123, Digits: 3
Word: world, Digits: 0

*/

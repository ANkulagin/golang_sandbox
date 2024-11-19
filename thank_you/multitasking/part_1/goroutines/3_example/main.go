package main

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc возвращает следующее слово из генератора.
type nextFunc func() string

// counter хранит количество цифр в каждом слове.
// Ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next().
func countDigitsInWords(next nextFunc) counter {
	stats := counter{}

	for {
		word := next()
		if word == "" {
			break
		}
		count := countDigits(word)
		stats[word] = count
	}

	return stats
}

// countDigits возвращает количество цифр в строке.
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats печатает количество цифр в словах.
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator возвращает генератор,
// который выдает слова из фразы.
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}

// Output:
/*
0ne: 1
1wo: 1
thr33: 2
4068: 4
*/

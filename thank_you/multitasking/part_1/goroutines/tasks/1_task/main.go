package main

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

func main() {
	input := "Hello 123 world 4567"
	words := strings.Fields(input)

	var wg sync.WaitGroup
	var syncStats sync.Map

	// начало решения
	for i := 0; i < len(words); i++ {
		wg.Add(1)
		go func(word string) {
			defer wg.Done()
			count := countDigits(word)
			syncStats.Store(word, count)
		}(words[i])
	}
	// Посчитайте количество цифр в словах,
	// используя отдельную горутину для каждого слова.

	// Чтобы записать результаты подсчета,
	// используйте syncStats.Store(word, count)

	// В результате syncStats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	wg.Wait()

	fmt.Println("Все содержимое syncStats:")
	syncStats.Range(func(key, value interface{}) bool {
		fmt.Printf("Слово: %s, Количество цифр: %d\n", key, value)
		return true
	})
}

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
	// Отправьте слова в канал pending
	for {
		word := next()
		if word == "" {
			close(pending)
			break
		}
		pending <- word
	}
}

// Функция countWords читает слова из канала pending, подсчитывает цифры,
// отправляет результаты в канал counted, а по завершению своей работы сообщает о этом через канал done.
func countWords(pending <-chan string, counted chan<- pair, done chan struct{}) {
	for w := range pending {
		c := countDigits(w)
		counted <- pair{word: w, count: c}
	}
	// Когда канал pending закрыт и данные закончились, горутина завершает работу
	// и сообщает об этом в канал done.
	done <- struct{}{}
}

func fillStats(counted chan pair) counter {
	stats := make(counter)
	// Заполните stats из канала counted
	for result := range counted {
		stats[result.word] = result.count
	}
	return stats
}

// Основная управляющая функция, которая запускает все процессы:
// 1. Запускает отправку слов в канал pending.
// 2. Запускает 4 горутины для обработки слов в countWords.
// 3. Ожидает завершения всех 4 горутин и закрывает канал counted.
// 4. Собирает результаты и возвращает их.
func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	done := make(chan struct{})

	// Запускаем 4 горутины для обработки слов:
	for i := 0; i < 4; i++ {
		go countWords(pending, counted, done)
	}

	// Создадим анонимную горутину, которая будет ожидать, когда все 4 горутины завершат работу.
	// Получим 4 сигнала из канала done и после этого закроем канал counted.
	go func() {
		for i := 0; i < 4; i++ {
			<-done
		}
		close(counted)
	}()

	// Собираем статистику
	return fillStats(counted)
}

func main() {
	words := []string{"Hello", "123", "world", "4567"}
	index := 0

	next := func() string {
		if index >= len(words) {
			return ""
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

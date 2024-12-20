package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func generate(done <-chan struct{}) <-chan string {
	out := make(chan string)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	go func() {
		defer close(out)
		letters := "abcdefghijklmnopqrstuvwxyz"
		for {
			select {
			case <-done:
				return
			default:
				var sb strings.Builder
				for i := 0; i < 5; i++ {
					sb.WriteString(string(letters[r.Intn(len(letters))]))
				}
				word := sb.String()
				select {
				case <-done:
					return
				case out <- word:
				}
			}
		}
	}()
	return out
}

func hasUniqueLetters(word string) bool {
	seen := make(map[rune]struct{})
	for _, char := range word {
		if _, exists := seen[char]; exists {
			return false
		}
		seen[char] = struct{}{}
	}
	return true
}

func takeUnique(done <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		// начало решения
		// Фильтрация уникальных слов
		defer close(out)

		for {
			select {
			case <-done:
				return
			case word, ok := <-in:
				if !ok {
					return
				}
				if hasUniqueLetters(word) {
					select {
					case <-done:
						return
					case out <- word:
					}
				}
			}
		}
		// конец решения
	}()
	return out
}

func reverseWord(done <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		// начало решения
		// Переворот слов
		defer close(out)
		for {
			select {
			case <-done:
				return
			case word, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-done:
					return
				case out <- fmt.Sprintf("%s -> %s", word, reverse(word)):
				}
			}
		}
		// конец решения

	}()
	return out
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	n := 10
	done := make(chan struct{})
	defer close(done)

	c1 := generate(done)
	c2 := takeUnique(done, c1)
	c3 := reverseWord(done, c2)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		count := 0
		for word := range c3 {
			fmt.Println(word)
			count++
			if count == n {
				break
			}
		}
	}()

	wg.Wait()
}
